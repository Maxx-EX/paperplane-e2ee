#!/usr/bin/env python3
"""生成 FILE_MANIFEST.md / FILE_MANIFEST.json：所有文件路径 + 文件内容。

用法:
    python3 scripts/gen-manifest.py [--root .]

产物:
    FILE_MANIFEST.md    人类可读：目录树 + 每个文件的完整代码
    FILE_MANIFEST.json  机器可读：{files: [{path, bytes, sha256, language}]}
"""
from __future__ import annotations

import argparse
import hashlib
import json
import os
import sys
from pathlib import Path

SKIP_DIRS = {
    ".git", "build", ".dart_tool", ".gradle", ".kotlin", "node_modules",
    "__pycache__", ".idea", ".vscode", "Pods", ".packages",
}
SKIP_SUFFIX = {".zip", ".apk", ".aab", ".o", ".so", ".dylib", ".dll", ".lock"}
SKIP_FILES = {"FILE_MANIFEST.md", "FILE_MANIFEST.json", "FILE_MANIFEST.html"}

LANG = {
    ".dart": "dart", ".go": "go", ".kt": "kotlin", ".kts": "kotlin",
    ".c": "c", ".h": "c", ".cpp": "cpp", ".hpp": "cpp",
    ".sh": "bash", ".py": "python", ".yaml": "yaml", ".yml": "yaml",
    ".md": "markdown", ".html": "html", ".json": "json", ".sql": "sql",
    ".env": "dotenv", ".txt": "text", ".properties": "properties",
}


def language(path: Path) -> str:
    if path.name == "Dockerfile":
        return "dockerfile"
    if path.name == "Makefile":
        return "makefile"
    return LANG.get(path.suffix.lower(), "text")


BINARY_SUFFIX = {".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".bmp",
                 ".apk", ".aab", ".so", ".o", ".dylib", ".dll", ".zip",
                 ".gz", ".tar", ".pdf", ".db"}

def is_binary(p: Path) -> bool:
    return p.suffix.lower() in BINARY_SUFFIX

def walk(root: Path):
    for dirpath, dirnames, filenames in os.walk(root):
        dirnames[:] = sorted(d for d in dirnames if d not in SKIP_DIRS)
        for name in sorted(filenames):
            p = Path(dirpath) / name
            if name in SKIP_FILES or p.suffix in SKIP_SUFFIX:
                continue
            if p.is_symlink() or not p.is_file():
                continue
            yield p


def tree(root: Path) -> str:
    lines: list[str] = [root.name + "/"]
    entries = sorted(
        (p for p in root.rglob("*") if not any(part in SKIP_DIRS for part in p.parts)),
        key=lambda p: (len(p.relative_to(root).parts), str(p)),
    )
    seen: set[Path] = set()
    for p in entries:
        rel = p.relative_to(root)
        if rel.parts[0] in SKIP_DIRS or p.name in SKIP_FILES:
            continue
        for i in range(len(rel.parts) - 1):
            parent = Path(*rel.parts[: i + 1])
            if parent not in seen:
                seen.add(parent)
                lines.append("  " * (i + 1) + parent.name + "/")
        if p.is_file():
            lines.append("  " * len(rel.parts) + p.name)
    return "\n".join(lines)


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default=".")
    args = ap.parse_args()
    root = Path(args.root).resolve()

    files = list(walk(root))
    total = sum(p.stat().st_size for p in files)
    records = []
    for p in files:
        rel = p.relative_to(root).as_posix()
        data = p.read_bytes()
        records.append({
            "path": rel,
            "language": language(p),
            "bytes": len(data),
            "sha256": hashlib.sha256(data).hexdigest(),
            "lines": data.count(b"\n") + (0 if data.endswith(b"\n") else 1),
        })

    md: list[str] = []
    md.append("# FILE MANIFEST — Paperplane\n")
    md.append(
        "本文件由 `scripts/gen-manifest.py` 生成，收录仓库内每个文件的路径与完整代码。\n\n"
        f"- 文件总数：**{len(records)}**\n- 总体积：**{total / 1024:.1f} KiB**\n"
        f"- 生成时间：{os.environ.get('MANIFEST_TIME', 'build-time')}\n"
    )
    md.append("## 目录树\n\n```text\n" + tree(root) + "\n```\n")
    md.append("## 文件索引\n\n| # | 路径 | 语言 | 行数 | 大小 |\n|---|---|---|---|---|\n")
    for i, r in enumerate(records, 1):
        md.append(f"| {i} | `{r['path']}` | {r['language']} | {r['lines']} | {r['bytes']} B |\n")
    md.append("\n## 文件内容\n")
    for r in records:
        fpath = root / r["path"]
        if is_binary(fpath):
            md.append(f"\n### `{r['path']}`\n\n_二进制资源（{r['bytes']} B，sha256 `{r['sha256']}`）_。\n")
            continue
        md.append(f"\n### `{r['path']}`\n\n```{r['language']}\n")
        md.append(fpath.read_text(encoding="utf-8", errors="replace"))
        if not md[-1].endswith("\n"):
            md.append("\n")
        md.append("```\n")

    (root / "FILE_MANIFEST.md").write_text("".join(md), encoding="utf-8")
    (root / "FILE_MANIFEST.json").write_text(
        json.dumps({"project": "paperplane", "files": records}, ensure_ascii=False, indent=2),
        encoding="utf-8",
    )
    print(f"[gen-manifest] {len(records)} files, {total / 1024:.1f} KiB -> FILE_MANIFEST.md")
    return 0


if __name__ == "__main__":
    sys.exit(main())
