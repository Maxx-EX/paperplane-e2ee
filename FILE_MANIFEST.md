# FILE MANIFEST — Paperplane
本文件由 `scripts/gen-manifest.py` 生成，收录仓库内每个文件的路径与完整代码。

- 文件总数：**74**
- 总体积：**2183.6 KiB**
- 生成时间：2026-11-05
## 目录树

```text
paperplane-e2ee/
  .gitignore
  LICENSE
  Makefile
  README.md
  docker-compose.dev.yml
  client/
    analysis_options.yaml
    pubspec.yaml
  config/
    config.dev.yaml
    config.yaml
  deploy/
    .env
    Caddyfile
    caddy.env
    docker-compose.yml
  docs/
    ARCHITECTURE.md
    CRYPTO.md
    KNOWN_ERRORS.md
    SELF_HOSTING.md
    _config.yml
    index.html
  policy-sidecar/
    Dockerfile
    build.gradle.kts
    settings.gradle.kts
  scripts/
    gen-manifest.py
    render-config.sh
    termux-bootstrap.sh
  server/
    Dockerfile
    go.mod
    go.sum
  xdp/
    Makefile
    config.env
    android/
    assets/
      app_icon.png
      icon-512.png
    lib/
      main.dart
    web/
      index.html
    crowdsec/
      acquis.yaml
      config.yaml.local
    assets/
      logo.png
    config/
      application.yaml
    src/
    cmd/
    internal/
    migrations/
      0001_init.sql
    src/
      xdp_filter.c
      xdp_loader.cpp
      app/
      config/
        app_config.dart
      crypto/
        password.dart
        signal_service.dart
      net/
        api_client.dart
        realtime.dart
      state/
        chat_store.dart
      ui/
        auth_page.dart
        chat_page.dart
        theme.dart
      main/
      server/
        main.go
      auth/
        token.go
      config/
        config.go
      defense/
        bans.go
        limits.go
      httpapi/
        handlers_auth.go
        handlers_envelopes.go
        handlers_keys.go
        middleware.go
        router.go
      store/
        postgres.go
        redis.go
      ws/
        hub.go
        src/
        widgets/
          composer.dart
          message_bubble.dart
        kotlin/
          main/
          io/
            res/
            paperplane/
              mipmap-hdpi/
                ic_launcher.png
              mipmap-mdpi/
                ic_launcher.png
              mipmap-xhdpi/
                ic_launcher.png
              mipmap-xxhdpi/
                ic_launcher.png
              mipmap-xxxhdpi/
                ic_launcher.png
              policy/
                BanPublisher.kt
                Config.kt
                Main.kt
                Rules.kt
```
## 文件索引

| # | 路径 | 语言 | 行数 | 大小 |
|---|---|---|---|---|
| 1 | `.gitignore` | text | 40 | 669 B |
| 2 | `LICENSE` | text | 661 | 34523 B |
| 3 | `Makefile` | makefile | 32 | 805 B |
| 4 | `README.md` | markdown | 93 | 4383 B |
| 5 | `docker-compose.dev.yml` | yaml | 37 | 1014 B |
| 6 | `client/analysis_options.yaml` | yaml | 5 | 98 B |
| 7 | `client/pubspec.yaml` | yaml | 39 | 893 B |
| 8 | `client/android/app/src/main/res/mipmap-hdpi/ic_launcher.png` | text | 11 | 3656 B |
| 9 | `client/android/app/src/main/res/mipmap-mdpi/ic_launcher.png` | text | 9 | 1837 B |
| 10 | `client/android/app/src/main/res/mipmap-xhdpi/ic_launcher.png` | text | 23 | 5924 B |
| 11 | `client/android/app/src/main/res/mipmap-xxhdpi/ic_launcher.png` | text | 57 | 12084 B |
| 12 | `client/android/app/src/main/res/mipmap-xxxhdpi/ic_launcher.png` | text | 81 | 20337 B |
| 13 | `client/assets/app_icon.png` | text | 3626 | 926208 B |
| 14 | `client/assets/icon-512.png` | text | 439 | 118572 B |
| 15 | `client/lib/main.dart` | dart | 43 | 1047 B |
| 16 | `client/lib/config/app_config.dart` | dart | 13 | 415 B |
| 17 | `client/lib/crypto/password.dart` | dart | 54 | 1960 B |
| 18 | `client/lib/crypto/signal_service.dart` | dart | 99 | 3752 B |
| 19 | `client/lib/net/api_client.dart` | dart | 130 | 3942 B |
| 20 | `client/lib/net/realtime.dart` | dart | 35 | 904 B |
| 21 | `client/lib/state/chat_store.dart` | dart | 161 | 5001 B |
| 22 | `client/lib/ui/auth_page.dart` | dart | 81 | 2624 B |
| 23 | `client/lib/ui/chat_page.dart` | dart | 78 | 2463 B |
| 24 | `client/lib/ui/theme.dart` | dart | 48 | 1614 B |
| 25 | `client/lib/ui/widgets/composer.dart` | dart | 60 | 1696 B |
| 26 | `client/lib/ui/widgets/message_bubble.dart` | dart | 56 | 1796 B |
| 27 | `client/web/index.html` | html | 15 | 479 B |
| 28 | `config/config.dev.yaml` | yaml | 61 | 1156 B |
| 29 | `config/config.yaml` | yaml | 79 | 2558 B |
| 30 | `deploy/.env` | text | 26 | 744 B |
| 31 | `deploy/Caddyfile` | text | 32 | 681 B |
| 32 | `deploy/caddy.env` | dotenv | 2 | 60 B |
| 33 | `deploy/docker-compose.yml` | yaml | 93 | 2842 B |
| 34 | `deploy/crowdsec/acquis.yaml` | yaml | 5 | 139 B |
| 35 | `deploy/crowdsec/config.yaml.local` | text | 17 | 508 B |
| 36 | `docs/ARCHITECTURE.md` | markdown | 50 | 1653 B |
| 37 | `docs/CRYPTO.md` | markdown | 65 | 2321 B |
| 38 | `docs/KNOWN_ERRORS.md` | markdown | 26 | 2714 B |
| 39 | `docs/SELF_HOSTING.md` | markdown | 70 | 2081 B |
| 40 | `docs/_config.yml` | yaml | 7 | 191 B |
| 41 | `docs/index.html` | html | 308 | 11486 B |
| 42 | `docs/assets/logo.png` | text | 3626 | 926208 B |
| 43 | `policy-sidecar/Dockerfile` | dockerfile | 11 | 329 B |
| 44 | `policy-sidecar/build.gradle.kts` | kotlin | 28 | 671 B |
| 45 | `policy-sidecar/settings.gradle.kts` | kotlin | 1 | 39 B |
| 46 | `policy-sidecar/config/application.yaml` | yaml | 22 | 655 B |
| 47 | `policy-sidecar/src/main/kotlin/io/paperplane/policy/BanPublisher.kt` | kotlin | 82 | 3146 B |
| 48 | `policy-sidecar/src/main/kotlin/io/paperplane/policy/Config.kt` | kotlin | 68 | 2849 B |
| 49 | `policy-sidecar/src/main/kotlin/io/paperplane/policy/Main.kt` | kotlin | 73 | 2485 B |
| 50 | `policy-sidecar/src/main/kotlin/io/paperplane/policy/Rules.kt` | kotlin | 67 | 2528 B |
| 51 | `scripts/gen-manifest.py` | python | 137 | 4955 B |
| 52 | `scripts/render-config.sh` | bash | 236 | 8475 B |
| 53 | `scripts/termux-bootstrap.sh` | bash | 28 | 947 B |
| 54 | `server/Dockerfile` | dockerfile | 17 | 542 B |
| 55 | `server/go.mod` | text | 23 | 730 B |
| 56 | `server/go.sum` | text | 49 | 4344 B |
| 57 | `server/cmd/server/main.go` | go | 229 | 6403 B |
| 58 | `server/internal/auth/token.go` | go | 206 | 6003 B |
| 59 | `server/internal/config/config.go` | go | 308 | 9606 B |
| 60 | `server/internal/defense/bans.go` | go | 79 | 1787 B |
| 61 | `server/internal/defense/limits.go` | go | 159 | 4737 B |
| 62 | `server/internal/httpapi/handlers_auth.go` | go | 284 | 9598 B |
| 63 | `server/internal/httpapi/handlers_envelopes.go` | go | 341 | 10667 B |
| 64 | `server/internal/httpapi/handlers_keys.go` | go | 222 | 6563 B |
| 65 | `server/internal/httpapi/middleware.go` | go | 125 | 3879 B |
| 66 | `server/internal/httpapi/router.go` | go | 199 | 6199 B |
| 67 | `server/internal/store/postgres.go` | go | 295 | 9531 B |
| 68 | `server/internal/store/redis.go` | go | 31 | 806 B |
| 69 | `server/internal/ws/hub.go` | go | 199 | 4448 B |
| 70 | `server/migrations/0001_init.sql` | sql | 43 | 1474 B |
| 71 | `xdp/Makefile` | makefile | 26 | 719 B |
| 72 | `xdp/config.env` | dotenv | 4 | 89 B |
| 73 | `xdp/src/xdp_filter.c` | c | 68 | 2016 B |
| 74 | `xdp/src/xdp_loader.cpp` | cpp | 147 | 4768 B |

## 文件内容

### `.gitignore`

```text
# ---- secrets & generated ----
deploy/.env
deploy/caddy.env
deploy/crowdsec/acquis.yaml
policy-sidecar/config/application.yaml
xdp/config.env
*.zip

# ---- Go ----
server/paperplane-server
server/**/*.test
server/coverage.out

# ---- Flutter / Dart ----
client/.dart_tool/
client/.packages
client/pubspec.lock
client/build/
client/.flutter-plugins
client/.flutter-plugins-dependencies
client/android/.gradle/
client/ios/Pods/
client/web/.dart_tool/

# ---- Kotlin / Gradle ----
policy-sidecar/.gradle/
policy-sidecar/build/
policy-sidecar/.kotlin/

# ---- XDP build ----
xdp/build/
xdp/*.o
xdp/vmlinux.h

# ---- editors / OS ----
.idea/
.vscode/
.DS_Store
*.swp
*.log
```

### `LICENSE`

```text
                    GNU AFFERO GENERAL PUBLIC LICENSE
                       Version 3, 19 November 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.

                            Preamble

  The GNU Affero General Public License is a free, copyleft license for
software and other kinds of works, specifically designed to ensure
cooperation with the community in the case of network server software.

  The licenses for most software and other practical works are designed
to take away your freedom to share and change the works.  By contrast,
our General Public Licenses are intended to guarantee your freedom to
share and change all versions of a program--to make sure it remains free
software for all its users.

  When we speak of free software, we are referring to freedom, not
price.  Our General Public Licenses are designed to make sure that you
have the freedom to distribute copies of free software (and charge for
them if you wish), that you receive source code or can get it if you
want it, that you can change the software or use pieces of it in new
free programs, and that you know you can do these things.

  Developers that use our General Public Licenses protect your rights
with two steps: (1) assert copyright on the software, and (2) offer
you this License which gives you legal permission to copy, distribute
and/or modify the software.

  A secondary benefit of defending all users' freedom is that
improvements made in alternate versions of the program, if they
receive widespread use, become available for other developers to
incorporate.  Many developers of free software are heartened and
encouraged by the resulting cooperation.  However, in the case of
software used on network servers, this result may fail to come about.
The GNU General Public License permits making a modified version and
letting the public access it on a server without ever releasing its
source code to the public.

  The GNU Affero General Public License is designed specifically to
ensure that, in such cases, the modified source code becomes available
to the community.  It requires the operator of a network server to
provide the source code of the modified version running there to the
users of that server.  Therefore, public use of a modified version, on
a publicly accessible server, gives the public access to the source
code of the modified version.

  An older license, called the Affero General Public License and
published by Affero, was designed to accomplish similar goals.  This is
a different license, not a version of the Affero GPL, but Affero has
released a new version of the Affero GPL which permits relicensing under
this license.

  The precise terms and conditions for copying, distribution and
modification follow.

                       TERMS AND CONDITIONS

  0. Definitions.

  "This License" refers to version 3 of the GNU Affero General Public License.

  "Copyright" also means copyright-like laws that apply to other kinds of
works, such as semiconductor masks.

  "The Program" refers to any copyrightable work licensed under this
License.  Each licensee is addressed as "you".  "Licensees" and
"recipients" may be individuals or organizations.

  To "modify" a work means to copy from or adapt all or part of the work
in a fashion requiring copyright permission, other than the making of an
exact copy.  The resulting work is called a "modified version" of the
earlier work or a work "based on" the earlier work.

  A "covered work" means either the unmodified Program or a work based
on the Program.

  To "propagate" a work means to do anything with it that, without
permission, would make you directly or secondarily liable for
infringement under applicable copyright law, except executing it on a
computer or modifying a private copy.  Propagation includes copying,
distribution (with or without modification), making available to the
public, and in some countries other activities as well.

  To "convey" a work means any kind of propagation that enables other
parties to make or receive copies.  Mere interaction with a user through
a computer network, with no transfer of a copy, is not conveying.

  An interactive user interface displays "Appropriate Legal Notices"
to the extent that it includes a convenient and prominently visible
feature that (1) displays an appropriate copyright notice, and (2)
tells the user that there is no warranty for the work (except to the
extent that warranties are provided), that licensees may convey the
work under this License, and how to view a copy of this License.  If
the interface presents a list of user commands or options, such as a
menu, a prominent item in the list meets this criterion.

  1. Source Code.

  The "source code" for a work means the preferred form of the work
for making modifications to it.  "Object code" means any non-source
form of a work.

  A "Standard Interface" means an interface that either is an official
standard defined by a recognized standards body, or, in the case of
interfaces specified for a particular programming language, one that
is widely used among developers working in that language.

  The "System Libraries" of an executable work include anything, other
than the work as a whole, that (a) is included in the normal form of
packaging a Major Component, but which is not part of that Major
Component, and (b) serves only to enable use of the work with that
Major Component, or to implement a Standard Interface for which an
implementation is available to the public in source code form.  A
"Major Component", in this context, means a major essential component
(kernel, window system, and so on) of the specific operating system
(if any) on which the executable work runs, or a compiler used to
produce the work, or an object code interpreter used to run it.

  The "Corresponding Source" for a work in object code form means all
the source code needed to generate, install, and (for an executable
work) run the object code and to modify the work, including scripts to
control those activities.  However, it does not include the work's
System Libraries, or general-purpose tools or generally available free
programs which are used unmodified in performing those activities but
which are not part of the work.  For example, Corresponding Source
includes interface definition files associated with source files for
the work, and the source code for shared libraries and dynamically
linked subprograms that the work is specifically designed to require,
such as by intimate data communication or control flow between those
subprograms and other parts of the work.

  The Corresponding Source need not include anything that users
can regenerate automatically from other parts of the Corresponding
Source.

  The Corresponding Source for a work in source code form is that
same work.

  2. Basic Permissions.

  All rights granted under this License are granted for the term of
copyright on the Program, and are irrevocable provided the stated
conditions are met.  This License explicitly affirms your unlimited
permission to run the unmodified Program.  The output from running a
covered work is covered by this License only if the output, given its
content, constitutes a covered work.  This License acknowledges your
rights of fair use or other equivalent, as provided by copyright law.

  You may make, run and propagate covered works that you do not
convey, without conditions so long as your license otherwise remains
in force.  You may convey covered works to others for the sole purpose
of having them make modifications exclusively for you, or provide you
with facilities for running those works, provided that you comply with
the terms of this License in conveying all material for which you do
not control copyright.  Those thus making or running the covered works
for you must do so exclusively on your behalf, under your direction
and control, on terms that prohibit them from making any copies of
your copyrighted material outside their relationship with you.

  Conveying under any other circumstances is permitted solely under
the conditions stated below.  Sublicensing is not allowed; section 10
makes it unnecessary.

  3. Protecting Users' Legal Rights From Anti-Circumvention Law.

  No covered work shall be deemed part of an effective technological
measure under any applicable law fulfilling obligations under article
11 of the WIPO copyright treaty adopted on 20 December 1996, or
similar laws prohibiting or restricting circumvention of such
measures.

  When you convey a covered work, you waive any legal power to forbid
circumvention of technological measures to the extent such circumvention
is effected by exercising rights under this License with respect to
the covered work, and you disclaim any intention to limit operation or
modification of the work as a means of enforcing, against the work's
users, your or third parties' legal rights to forbid circumvention of
technological measures.

  4. Conveying Verbatim Copies.

  You may convey verbatim copies of the Program's source code as you
receive it, in any medium, provided that you conspicuously and
appropriately publish on each copy an appropriate copyright notice;
keep intact all notices stating that this License and any
non-permissive terms added in accord with section 7 apply to the code;
keep intact all notices of the absence of any warranty; and give all
recipients a copy of this License along with the Program.

  You may charge any price or no price for each copy that you convey,
and you may offer support or warranty protection for a fee.

  5. Conveying Modified Source Versions.

  You may convey a work based on the Program, or the modifications to
produce it from the Program, in the form of source code under the
terms of section 4, provided that you also meet all of these conditions:

    a) The work must carry prominent notices stating that you modified
    it, and giving a relevant date.

    b) The work must carry prominent notices stating that it is
    released under this License and any conditions added under section
    7.  This requirement modifies the requirement in section 4 to
    "keep intact all notices".

    c) You must license the entire work, as a whole, under this
    License to anyone who comes into possession of a copy.  This
    License will therefore apply, along with any applicable section 7
    additional terms, to the whole of the work, and all its parts,
    regardless of how they are packaged.  This License gives no
    permission to license the work in any other way, but it does not
    invalidate such permission if you have separately received it.

    d) If the work has interactive user interfaces, each must display
    Appropriate Legal Notices; however, if the Program has interactive
    interfaces that do not display Appropriate Legal Notices, your
    work need not make them do so.

  A compilation of a covered work with other separate and independent
works, which are not by their nature extensions of the covered work,
and which are not combined with it such as to form a larger program,
in or on a volume of a storage or distribution medium, is called an
"aggregate" if the compilation and its resulting copyright are not
used to limit the access or legal rights of the compilation's users
beyond what the individual works permit.  Inclusion of a covered work
in an aggregate does not cause this License to apply to the other
parts of the aggregate.

  6. Conveying Non-Source Forms.

  You may convey a covered work in object code form under the terms
of sections 4 and 5, provided that you also convey the
machine-readable Corresponding Source under the terms of this License,
in one of these ways:

    a) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by the
    Corresponding Source fixed on a durable physical medium
    customarily used for software interchange.

    b) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by a
    written offer, valid for at least three years and valid for as
    long as you offer spare parts or customer support for that product
    model, to give anyone who possesses the object code either (1) a
    copy of the Corresponding Source for all the software in the
    product that is covered by this License, on a durable physical
    medium customarily used for software interchange, for a price no
    more than your reasonable cost of physically performing this
    conveying of source, or (2) access to copy the
    Corresponding Source from a network server at no charge.

    c) Convey individual copies of the object code with a copy of the
    written offer to provide the Corresponding Source.  This
    alternative is allowed only occasionally and noncommercially, and
    only if you received the object code with such an offer, in accord
    with subsection 6b.

    d) Convey the object code by offering access from a designated
    place (gratis or for a charge), and offer equivalent access to the
    Corresponding Source in the same way through the same place at no
    further charge.  You need not require recipients to copy the
    Corresponding Source along with the object code.  If the place to
    copy the object code is a network server, the Corresponding Source
    may be on a different server (operated by you or a third party)
    that supports equivalent copying facilities, provided you maintain
    clear directions next to the object code saying where to find the
    Corresponding Source.  Regardless of what server hosts the
    Corresponding Source, you remain obligated to ensure that it is
    available for as long as needed to satisfy these requirements.

    e) Convey the object code using peer-to-peer transmission, provided
    you inform other peers where the object code and Corresponding
    Source of the work are being offered to the general public at no
    charge under subsection 6d.

  A separable portion of the object code, whose source code is excluded
from the Corresponding Source as a System Library, need not be
included in conveying the object code work.

  A "User Product" is either (1) a "consumer product", which means any
tangible personal property which is normally used for personal, family,
or household purposes, or (2) anything designed or sold for incorporation
into a dwelling.  In determining whether a product is a consumer product,
doubtful cases shall be resolved in favor of coverage.  For a particular
product received by a particular user, "normally used" refers to a
typical or common use of that class of product, regardless of the status
of the particular user or of the way in which the particular user
actually uses, or expects or is expected to use, the product.  A product
is a consumer product regardless of whether the product has substantial
commercial, industrial or non-consumer uses, unless such uses represent
the only significant mode of use of the product.

  "Installation Information" for a User Product means any methods,
procedures, authorization keys, or other information required to install
and execute modified versions of a covered work in that User Product from
a modified version of its Corresponding Source.  The information must
suffice to ensure that the continued functioning of the modified object
code is in no case prevented or interfered with solely because
modification has been made.

  If you convey an object code work under this section in, or with, or
specifically for use in, a User Product, and the conveying occurs as
part of a transaction in which the right of possession and use of the
User Product is transferred to the recipient in perpetuity or for a
fixed term (regardless of how the transaction is characterized), the
Corresponding Source conveyed under this section must be accompanied
by the Installation Information.  But this requirement does not apply
if neither you nor any third party retains the ability to install
modified object code on the User Product (for example, the work has
been installed in ROM).

  The requirement to provide Installation Information does not include a
requirement to continue to provide support service, warranty, or updates
for a work that has been modified or installed by the recipient, or for
the User Product in which it has been modified or installed.  Access to a
network may be denied when the modification itself materially and
adversely affects the operation of the network or violates the rules and
protocols for communication across the network.

  Corresponding Source conveyed, and Installation Information provided,
in accord with this section must be in a format that is publicly
documented (and with an implementation available to the public in
source code form), and must require no special password or key for
unpacking, reading or copying.

  7. Additional Terms.

  "Additional permissions" are terms that supplement the terms of this
License by making exceptions from one or more of its conditions.
Additional permissions that are applicable to the entire Program shall
be treated as though they were included in this License, to the extent
that they are valid under applicable law.  If additional permissions
apply only to part of the Program, that part may be used separately
under those permissions, but the entire Program remains governed by
this License without regard to the additional permissions.

  When you convey a copy of a covered work, you may at your option
remove any additional permissions from that copy, or from any part of
it.  (Additional permissions may be written to require their own
removal in certain cases when you modify the work.)  You may place
additional permissions on material, added by you to a covered work,
for which you have or can give appropriate copyright permission.

  Notwithstanding any other provision of this License, for material you
add to a covered work, you may (if authorized by the copyright holders of
that material) supplement the terms of this License with terms:

    a) Disclaiming warranty or limiting liability differently from the
    terms of sections 15 and 16 of this License; or

    b) Requiring preservation of specified reasonable legal notices or
    author attributions in that material or in the Appropriate Legal
    Notices displayed by works containing it; or

    c) Prohibiting misrepresentation of the origin of that material, or
    requiring that modified versions of such material be marked in
    reasonable ways as different from the original version; or

    d) Limiting the use for publicity purposes of names of licensors or
    authors of the material; or

    e) Declining to grant rights under trademark law for use of some
    trade names, trademarks, or service marks; or

    f) Requiring indemnification of licensors and authors of that
    material by anyone who conveys the material (or modified versions of
    it) with contractual assumptions of liability to the recipient, for
    any liability that these contractual assumptions directly impose on
    those licensors and authors.

  All other non-permissive additional terms are considered "further
restrictions" within the meaning of section 10.  If the Program as you
received it, or any part of it, contains a notice stating that it is
governed by this License along with a term that is a further
restriction, you may remove that term.  If a license document contains
a further restriction but permits relicensing or conveying under this
License, you may add to a covered work material governed by the terms
of that license document, provided that the further restriction does
not survive such relicensing or conveying.

  If you add terms to a covered work in accord with this section, you
must place, in the relevant source files, a statement of the
additional terms that apply to those files, or a notice indicating
where to find the applicable terms.

  Additional terms, permissive or non-permissive, may be stated in the
form of a separately written license, or stated as exceptions;
the above requirements apply either way.

  8. Termination.

  You may not propagate or modify a covered work except as expressly
provided under this License.  Any attempt otherwise to propagate or
modify it is void, and will automatically terminate your rights under
this License (including any patent licenses granted under the third
paragraph of section 11).

  However, if you cease all violation of this License, then your
license from a particular copyright holder is reinstated (a)
provisionally, unless and until the copyright holder explicitly and
finally terminates your license, and (b) permanently, if the copyright
holder fails to notify you of the violation by some reasonable means
prior to 60 days after the cessation.

  Moreover, your license from a particular copyright holder is
reinstated permanently if the copyright holder notifies you of the
violation by some reasonable means, this is the first time you have
received notice of violation of this License (for any work) from that
copyright holder, and you cure the violation prior to 30 days after
your receipt of the notice.

  Termination of your rights under this section does not terminate the
licenses of parties who have received copies or rights from you under
this License.  If your rights have been terminated and not permanently
reinstated, you do not qualify to receive new licenses for the same
material under section 10.

  9. Acceptance Not Required for Having Copies.

  You are not required to accept this License in order to receive or
run a copy of the Program.  Ancillary propagation of a covered work
occurring solely as a consequence of using peer-to-peer transmission
to receive a copy likewise does not require acceptance.  However,
nothing other than this License grants you permission to propagate or
modify any covered work.  These actions infringe copyright if you do
not accept this License.  Therefore, by modifying or propagating a
covered work, you indicate your acceptance of this License to do so.

  10. Automatic Licensing of Downstream Recipients.

  Each time you convey a covered work, the recipient automatically
receives a license from the original licensors, to run, modify and
propagate that work, subject to this License.  You are not responsible
for enforcing compliance by third parties with this License.

  An "entity transaction" is a transaction transferring control of an
organization, or substantially all assets of one, or subdividing an
organization, or merging organizations.  If propagation of a covered
work results from an entity transaction, each party to that
transaction who receives a copy of the work also receives whatever
licenses to the work the party's predecessor in interest had or could
give under the previous paragraph, plus a right to possession of the
Corresponding Source of the work from the predecessor in interest, if
the predecessor has it or can get it with reasonable efforts.

  You may not impose any further restrictions on the exercise of the
rights granted or affirmed under this License.  For example, you may
not impose a license fee, royalty, or other charge for exercise of
rights granted under this License, and you may not initiate litigation
(including a cross-claim or counterclaim in a lawsuit) alleging that
any patent claim is infringed by making, using, selling, offering for
sale, or importing the Program or any portion of it.

  11. Patents.

  A "contributor" is a copyright holder who authorizes use under this
License of the Program or a work on which the Program is based.  The
work thus licensed is called the contributor's "contributor version".

  A contributor's "essential patent claims" are all patent claims
owned or controlled by the contributor, whether already acquired or
hereafter acquired, that would be infringed by some manner, permitted
by this License, of making, using, or selling its contributor version,
but do not include claims that would be infringed only as a
consequence of further modification of the contributor version.  For
purposes of this definition, "control" includes the right to grant
patent sublicenses in a manner consistent with the requirements of
this License.

  Each contributor grants you a non-exclusive, worldwide, royalty-free
patent license under the contributor's essential patent claims, to
make, use, sell, offer for sale, import and otherwise run, modify and
propagate the contents of its contributor version.

  In the following three paragraphs, a "patent license" is any express
agreement or commitment, however denominated, not to enforce a patent
(such as an express permission to practice a patent or covenant not to
sue for patent infringement).  To "grant" such a patent license to a
party means to make such an agreement or commitment not to enforce a
patent against the party.

  If you convey a covered work, knowingly relying on a patent license,
and the Corresponding Source of the work is not available for anyone
to copy, free of charge and under the terms of this License, through a
publicly available network server or other readily accessible means,
then you must either (1) cause the Corresponding Source to be so
available, or (2) arrange to deprive yourself of the benefit of the
patent license for this particular work, or (3) arrange, in a manner
consistent with the requirements of this License, to extend the patent
license to downstream recipients.  "Knowingly relying" means you have
actual knowledge that, but for the patent license, your conveying the
covered work in a country, or your recipient's use of the covered work
in a country, would infringe one or more identifiable patents in that
country that you have reason to believe are valid.

  If, pursuant to or in connection with a single transaction or
arrangement, you convey, or propagate by procuring conveyance of, a
covered work, and grant a patent license to some of the parties
receiving the covered work authorizing them to use, propagate, modify
or convey a specific copy of the covered work, then the patent license
you grant is automatically extended to all recipients of the covered
work and works based on it.

  A patent license is "discriminatory" if it does not include within
the scope of its coverage, prohibits the exercise of, or is
conditioned on the non-exercise of one or more of the rights that are
specifically granted under this License.  You may not convey a covered
work if you are a party to an arrangement with a third party that is
in the business of distributing software, under which you make payment
to the third party based on the extent of your activity of conveying
the work, and under which the third party grants, to any of the
parties who would receive the covered work from you, a discriminatory
patent license (a) in connection with copies of the covered work
conveyed by you (or copies made from those copies), or (b) primarily
for and in connection with specific products or compilations that
contain the covered work, unless you entered into that arrangement,
or that patent license was granted, prior to 28 March 2007.

  Nothing in this License shall be construed as excluding or limiting
any implied license or other defenses to infringement that may
otherwise be available to you under applicable patent law.

  12. No Surrender of Others' Freedom.

  If conditions are imposed on you (whether by court order, agreement or
otherwise) that contradict the conditions of this License, they do not
excuse you from the conditions of this License.  If you cannot convey a
covered work so as to satisfy simultaneously your obligations under this
License and any other pertinent obligations, then as a consequence you may
not convey it at all.  For example, if you agree to terms that obligate you
to collect a royalty for further conveying from those to whom you convey
the Program, the only way you could satisfy both those terms and this
License would be to refrain entirely from conveying the Program.

  13. Remote Network Interaction; Use with the GNU General Public License.

  Notwithstanding any other provision of this License, if you modify the
Program, your modified version must prominently offer all users
interacting with it remotely through a computer network (if your version
supports such interaction) an opportunity to receive the Corresponding
Source of your version by providing access to the Corresponding Source
from a network server at no charge, through some standard or customary
means of facilitating copying of software.  This Corresponding Source
shall include the Corresponding Source for any work covered by version 3
of the GNU General Public License that is incorporated pursuant to the
following paragraph.

  Notwithstanding any other provision of this License, you have
permission to link or combine any covered work with a work licensed
under version 3 of the GNU General Public License into a single
combined work, and to convey the resulting work.  The terms of this
License will continue to apply to the part which is the covered work,
but the work with which it is combined will remain governed by version
3 of the GNU General Public License.

  14. Revised Versions of this License.

  The Free Software Foundation may publish revised and/or new versions of
the GNU Affero General Public License from time to time.  Such new versions
will be similar in spirit to the present version, but may differ in detail to
address new problems or concerns.

  Each version is given a distinguishing version number.  If the
Program specifies that a certain numbered version of the GNU Affero General
Public License "or any later version" applies to it, you have the
option of following the terms and conditions either of that numbered
version or of any later version published by the Free Software
Foundation.  If the Program does not specify a version number of the
GNU Affero General Public License, you may choose any version ever published
by the Free Software Foundation.

  If the Program specifies that a proxy can decide which future
versions of the GNU Affero General Public License can be used, that proxy's
public statement of acceptance of a version permanently authorizes you
to choose that version for the Program.

  Later license versions may give you additional or different
permissions.  However, no additional obligations are imposed on any
author or copyright holder as a result of your choosing to follow a
later version.

  15. Disclaimer of Warranty.

  THERE IS NO WARRANTY FOR THE PROGRAM, TO THE EXTENT PERMITTED BY
APPLICABLE LAW.  EXCEPT WHEN OTHERWISE STATED IN WRITING THE COPYRIGHT
HOLDERS AND/OR OTHER PARTIES PROVIDE THE PROGRAM "AS IS" WITHOUT WARRANTY
OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING, BUT NOT LIMITED TO,
THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
PURPOSE.  THE ENTIRE RISK AS TO THE QUALITY AND PERFORMANCE OF THE PROGRAM
IS WITH YOU.  SHOULD THE PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF
ALL NECESSARY SERVICING, REPAIR OR CORRECTION.

  16. Limitation of Liability.

  IN NO EVENT UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING
WILL ANY COPYRIGHT HOLDER, OR ANY OTHER PARTY WHO MODIFIES AND/OR CONVEYS
THE PROGRAM AS PERMITTED ABOVE, BE LIABLE TO YOU FOR DAMAGES, INCLUDING ANY
GENERAL, SPECIAL, INCIDENTAL OR CONSEQUENTIAL DAMAGES ARISING OUT OF THE
USE OR INABILITY TO USE THE PROGRAM (INCLUDING BUT NOT LIMITED TO LOSS OF
DATA OR DATA BEING RENDERED INACCURATE OR LOSSES SUSTAINED BY YOU OR THIRD
PARTIES OR A FAILURE OF THE PROGRAM TO OPERATE WITH ANY OTHER PROGRAMS),
EVEN IF SUCH HOLDER OR OTHER PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF
SUCH DAMAGES.

  17. Interpretation of Sections 15 and 16.

  If the disclaimer of warranty and limitation of liability provided
above cannot be given local legal effect according to their terms,
reviewing courts shall apply local law that most closely approximates
an absolute waiver of all civil liability in connection with the
Program, unless a warranty or assumption of liability accompanies a
copy of the Program in return for a fee.

                     END OF TERMS AND CONDITIONS

            How to Apply These Terms to Your New Programs

  If you develop a new program, and you want it to be of the greatest
possible use to the public, the best way to achieve this is to make it
free software which everyone can redistribute and change under these terms.

  To do so, attach the following notices to the program.  It is safest
to attach them to the start of each source file to most effectively
state the exclusion of warranty; and each file should have at least
the "copyright" line and a pointer to where the full notice is found.

    <one line to give the program's name and a brief idea of what it does.>
    Copyright (C) <year>  <name of author>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

Also add information on how to contact you by electronic and paper mail.

  If your software can interact with users remotely through a computer
network, you should also make sure that it provides a way for users to
get its source.  For example, if your program is a web application, its
interface could display a "Source" link that leads users to an archive
of the code.  There are many ways you could offer source, and different
solutions will be better for different programs; see section 13 for the
specific requirements.

  You should also get your employer (if you work as a programmer) or school,
if any, to sign a "copyright disclaimer" for the program, if necessary.
For more information on this, and how to apply and follow the GNU AGPL, see
<https://www.gnu.org/licenses/>.
```

### `Makefile`

```makefile
.PHONY: all config config-check build-server server dev client manifest release clean

all: build-server

config:
	./scripts/render-config.sh

config-check:
	./scripts/render-config.sh --check

build-server:
	cd server && CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o paperplane-server ./cmd/server

server:
	cd server && go run ./cmd/server

dev: config
	docker compose -f docker-compose.dev.yml up -d --build
	@echo "Go 服务端已起；客户端： cd client && flutter pub get && flutter run -d chrome"

client:
	cd client && flutter pub get && flutter run -d chrome

manifest:
	python3 scripts/gen-manifest.py --root .

release: config-check build-server manifest
	@mkdir -p dist
	@echo "产物：server/paperplane-server, FILE_MANIFEST.md"

clean:
	rm -f server/paperplane-server dist/*.zip
```

### `README.md`

```markdown
# Paperplane

**纸飞机风格的零知识 E2EE 加密聊天系统** — 服务端不解密、不记录密文、不存储可逆账号字段，私钥仅客户端安全存储。

- 客户端：Flutter 3.x（`flutter_chat_ui`）+ `libsignal_protocol_dart`（X3DH + Double Ratchet）
- 服务端：Go 1.25（net/http + gorilla/websocket）+ PostgreSQL + Redis
- 防护：Kotlin sidecar（协程 + Redis Streams）、C/C++ eBPF/XDP（libbpf LPM 黑名单）、CrowdSec IPS、Caddy 终止 TLS
- 许可证：**AGPL-3.0**，开源仓库

```
┌──────────┐  注册/登录  ┌──────────┐   转发密文信封   ┌────────────┐
│  Flutter │ ──────────▶ │   Go     │ ◀─────────────▶ │  PostgreSQL│
│  客户端  │             │  Relay   │                  │   envelopes│
│ (libsignal)│ ◀────────── │ (零知识) │  不存私钥/明文   └────────────┘
└──────────┘  X3DH/DR    └────┬─────┘
                              │ 事件流
                              ▼
                        ┌──────────┐   ┌──────────┐
                        │ Kotlin  │──▶│ CrowdSec │
                        │ sidecar │   └──────────┘
                        └────┬─────┘
                             │ ban:ip ZSET
                             ▼
                        ┌──────────┐
                        │ XDP/TC  │ 内核态丢包
                        │  eBPF    │
                        └──────────┘
```

## 核心设计

- **账号单向链路**：
  `account_id = SHA256(username ‖ pepper)`；
  `argon2id = Argon2id(password, salt=account_id[:16], m=64MiB, t=3, p=1)`
  服务端只存 `account_id + PHC`，登录用 `HMAC-SHA256(credential, nonce)` 挑战响应，密码零接触。
- **E2EE**：首次登录生成 IdentityKey / SignedPreKey / OneTimePreKey / KyberPreKey（占位），公钥上传，私钥入 `flutter_secure_storage`；`SessionBuilder.processPreKeyBundle` 完成 X3DH，`SessionCipher` 收发。
- **零知识**：服务端只转发密文与信封，投递后清除路由元数据，支持 Sealed Sender（盲令牌）。
- **单一配置源**：一切服务器配置收敛到 `config/config.yaml`，`make config` 派生 `.env` / Caddy / CrowdSec / Kotlin / XDP。
- **UI**：Material 3 深/浅双主题，蓝灰主色；渐变气泡 18px 圆角、单/双勾已读、文本/图片/文件输入栏。

## 快速开始

```bash
# 1. 配置
cp config/config.yaml config/config.yaml   # 改密钥（至少 32 字节）
make config-check                         # 校验字段

# 2. 本地闭环
docker compose -f docker-compose.dev.yml up -d
cd client && flutter pub get && flutter run -d chrome
# 在 Chrome 里注册两个账号即可互发密文
```

## 生产部署

```bash
./scripts/render-config.sh                 # 生成 deploy/.env 等
cd deploy && docker compose up -d
# DNS 指向本机后 Caddy 自动签 Let's Encrypt 证书
```

详见 [`docs/SELF_HOSTING.md`](docs/SELF_HOSTING.md)。

## 目录结构

```
config/config.yaml         # 唯一配置源
server/                    # Go 服务端（已通过 go build）
client/                    # Flutter 客户端
policy-sidecar/            # Kotlin 策略 sidecar
xdp/                      # C/C++ eBPF/XDP 内核过滤
deploy/                   # Caddy / CrowdSec / docker-compose
docs/                     # GitHub Pages 官网 + 文档
scripts/                  # 配置渲染 / 清单生成
```

## 防御层

| 层 | 技术 | 位置 |
|---|---|---|
| 协议层 | WS Origin 精确白名单 + IP 连接限流 | Go `CheckOrigin` |
| 速率层 | Redis Lua 令牌桶 + 目标滑动窗口 | Go `defense/limits.go` |
| 策略层 | Redis Streams 事件流 + 跨维规则 | Kotlin sidecar |
| 机器层 | CrowdSec LAPI 决策 | CrowdSec container |
| 内核层 | LPM_TRIE 黑名单 + SIGTERM detach | C/C++ XDP |

## 已知坑

见 [`docs/KNOWN_ERRORS.md`](docs/KNOWN_ERRORS.md)（A–J 协议层 + K-1–K-9 内核/策略层）。

## 许可证

AGPL-3.0 © Paperplane Contributors
```

### `docker-compose.dev.yml`

```yaml
# Paperplane 开发编排：仅 Postgres + Redis + Go 服务端
# 配合 `make dev`：docker compose 拉起后 flutter run -d chrome 直连
services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: paperplane
      POSTGRES_USER: paperplane
      POSTGRES_PASSWORD: dev-password
    ports:
      - "127.0.0.1:5432:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U paperplane"]
      interval: 3s
      retries: 30

  redis:
    image: redis:7-alpine
    command: ["redis-server", "--port", "6379"]
    ports:
      - "127.0.0.1:6379:6379"
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 3s
      retries: 30

  server:
    build: ./server
    depends_on:
      postgres: { condition: service_healthy }
      redis: { condition: service_healthy }
    network_mode: host
    volumes:
      - ./config/config.dev.yaml:/etc/paperplane/config.yaml:ro
    environment:
      CONFIG_PATH: /etc/paperplane/config.yaml
      MIGRATIONS_DIR: /migrations
```

### `client/analysis_options.yaml`

```yaml
include: package:flutter_lints/flutter.yaml

linter:
  rules:
    prefer_const_constructors: true
```

### `client/pubspec.yaml`

```yaml
name: paperplane
description: Zero-knowledge E2EE chat client (paperplane-style, AGPL-3.0)
publish_to: "none"
version: 0.1.0+1

environment:
  sdk: ">=3.5.0 <4.0.0"

dependencies:
  flutter:
    sdk: flutter
  flutter_chat_ui: 1.6.13
  flutter_chat_types: ^3.6.2
  libsignal_protocol_dart: ^0.8.2
  flutter_secure_storage: ^9.2.2
  serverpod_argon2: ^0.1.0
  http: ^1.2.2
  web_socket_channel: ^3.0.1
  crypto: ^3.0.5
  uuid: ^4.5.1
  provider: ^6.1.2

dev_dependencies:
  flutter_lints: ^5.0.0
  flutter_launcher_icons: ^0.14.1

flutter:
  uses-material-design: true
  assets:
    - assets/app_icon.png

# flutter_launcher_icons 配置：运行 `dart run flutter_launcher_icons` 生成 mipmap
flutter_launcher_icons:
  android: true
  ios: false
  image_path: "assets/app_icon.png"
  min_sdk_android: 24
  adaptive_icon_background: "#ffffff"
  adaptive_icon_foreground: "assets/app_icon.png"
```

### `client/android/app/src/main/res/mipmap-hdpi/ic_launcher.png`

_二进制资源（3656 B，sha256 `9ccb518dcf53ff0d7d7c351e4e83610c9053fa41f232de1968674e51636e8474`）_。

### `client/android/app/src/main/res/mipmap-mdpi/ic_launcher.png`

_二进制资源（1837 B，sha256 `6304e2bb01bcb814f35251b2401c8682bf2ee7f2859b21c33881463ff02af7cb`）_。

### `client/android/app/src/main/res/mipmap-xhdpi/ic_launcher.png`

_二进制资源（5924 B，sha256 `0a1aa679658abca235a0d2c7d6cf589e1784fd0e6462f79a7e453a611d098b18`）_。

### `client/android/app/src/main/res/mipmap-xxhdpi/ic_launcher.png`

_二进制资源（12084 B，sha256 `bab5309a2addaa1ea1b907c90ca859e830fe45c092edb15b4b6be7f5e5bdc434`）_。

### `client/android/app/src/main/res/mipmap-xxxhdpi/ic_launcher.png`

_二进制资源（20337 B，sha256 `f84c0bc0f1217e53c5ac92745ff0b16aa4db0117d45e83d829df7b00404afb0a`）_。

### `client/assets/app_icon.png`

_二进制资源（926208 B，sha256 `e2be9eac39cd665a690ad2597044f9f72d4e7014845697acf7625e023d2cb76f`）_。

### `client/assets/icon-512.png`

_二进制资源（118572 B，sha256 `94fbaca37b896bdf33dad5b17ff1872bdebddb931e9fcea8806d36e3bcde60dc`）_。

### `client/lib/main.dart`

```dart
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'state/chat_store.dart';
import 'ui/auth_page.dart';
import 'ui/chat_page.dart';
import 'ui/theme.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const PaperplaneApp());
}

class PaperplaneApp extends StatelessWidget {
  const PaperplaneApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => ChatStore()..bootstrap(),
      child: MaterialApp(
        title: 'Paperplane',
        debugShowCheckedModeBanner: false,
        theme: buildLightTheme(),
        darkTheme: buildDarkTheme(),
        themeMode: ThemeMode.system,
        home: const AuthGate(),
      ),
    );
  }
}

class AuthGate extends StatelessWidget {
  const AuthGate({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<ChatStore>(
      builder: (_, store, __) =>
          store.loggedIn ? const ChatPage() : const AuthPage(),
    );
  }
}
```

### `client/lib/config/app_config.dart`

```dart
// 服务端配置：通过 --dart-define=SERVER_URL=... 注入，默认连本机 8080
class AppConfig {
  static const String baseUrl = String.fromEnvironment(
    'SERVER_URL',
    defaultValue: 'http://127.0.0.1:8080',
  );

  static String get wsUrl {
    final u = Uri.parse(baseUrl);
    final scheme = u.scheme == 'https' ? 'wss' : 'ws';
    return '$scheme://${u.host}:${u.hasPort ? u.port : 80}/ws';
  }
}
```

### `client/lib/crypto/password.dart`

```dart
// 账号单向派生链：
//   account_id = SHA-256(username ‖ pepper)
//   hash       = Argon2id(password, salt = account_id[:16], m=64MiB, t=3, p=1)
//   proof      = HMAC-SHA256(credential, nonce)
//
// 服务端只存 account_id + PHC，从不接触明文口令。
import 'dart:convert';
import 'dart:isolate';

import 'package:crypto/crypto.dart';
import 'package:serverpod_argon2/serverpod_argon2.dart';

class Argon2Params {
  final int memoryKib;
  final int iterations;
  final int parallelism;
  const Argon2Params({
    required this.memoryKib,
    required this.iterations,
    required this.parallelism,
  });

  factory Argon2Params.fromServer(Map<String, dynamic> json) => Argon2Params(
        memoryKib: (json['argon2_memory_kib'] as num?)?.toInt() ?? 65536,
        iterations: (json['argon2_iterations'] as num?)?.toInt() ?? 3,
        parallelism: (json['argon2_parallelism'] as num?)?.toInt() ?? 1,
      );
}

class AccountPassword {
  static String accountId(String username, String pepper) =>
      sha256.convert(utf8.encode('${username.trim().toLowerCase()}$pepper')).toString();

  /// Argon2id 故意很慢：必须在 Isolate.run 里跑，绝不阻塞 UI。
  static Future<String> hashPassword(
    String password,
    String accountId,
    Argon2Params params,
  ) async {
    final argon2 = await Argon2.load();
    final salt = utf8.encode(accountId.substring(0, 16));
    // 注意：字段名随 serverpod_argon2 版本可能微调；如编译失败，
    // 仅需修改下面这一行（memory/iterations/parallelism）。
    final parameters = Argon2Parameters(
      memory: params.memoryKib,
      iterations: params.iterations,
      parallelism: params.parallelism,
    );
    return Isolate.run(() => argon2.hash(password, salt: salt, parameters: parameters));
  }

  static String proof(String credential, String nonce) =>
      Hmac(sha256, utf8.encode(credential)).convert(utf8.encode(nonce)).toString();
}
```

### `client/lib/crypto/signal_service.dart`

```dart
// Signal 协议层：X3DH + Double Ratchet 完全委托给 libsignal_protocol_dart。
// 本项目不自研任何 E2EE 算法。
// 注意：libsignal_protocol_dart 0.8.2 的部分构造器命名随版本微调，
// 如编译失败，仅需按官方 API 微调本文件中的字节反序列化调用。
import 'dart:convert';
import 'dart:typed_data';

import 'package:libsignal_protocol_dart/libsignal_protocol_dart.dart';

class SignalService {
  InMemorySignalProtocolStore? _store;
  static const int deviceId = 1;

  Future<void> ensureIdentity() async {
    if (_store != null) return;
    final identityKeyPair = generateIdentityKeyPair();
    final registrationId = generateRegistrationId(false);
    _store = InMemorySignalProtocolStore(identityKeyPair, registrationId);
    final spk = generateSignedPreKey(identityKeyPair, 0);
    await _store!.storeSignedPreKey(spk.id, spk);
    final preKeys = generatePreKeys(0, 100);
    for (final p in preKeys) {
      await _store!.storePreKey(p.id, p);
    }
  }

  /// 生成预密钥公钥包，只传公钥；私钥不出本机（内存即销毁）。
  Future<Map<String, dynamic>> buildBundle() async {
    await ensureIdentity();
    final s = _store!;
    final spk = await s.loadSignedPreKey(1);
    final ots = <Map<String, dynamic>>[];
    for (var id = 0; id < 100; id++) {
      try {
        final k = await s.loadPreKey(id);
        ots.add({'id': id, 'public': base64Encode(k.serialize())});
      } catch (_) {}
    }
    return {
      'identity_key': base64Encode(s.identityKeyPair().publicKey.serialize()),
      'signed_prekey': {
        'id': spk.id,
        'public': base64Encode(spk.keyPair.publicKey.serialize()),
        'signature': base64Encode(spk.signature),
      },
      'one_time_prekeys': ots,
    };
  }

  SignalProtocolAddress _addr(String remoteId) =>
      SignalProtocolAddress(remoteId, deviceId);

  /// 用服务端返回的预密钥包建立 X3DH 会话。
  Future<void> establishSession(
      String remoteId, Map<String, dynamic> bundle) async {
    await ensureIdentity();
    final builder = SessionBuilder(
        _store!, _store!, _store!, _store!, _addr(remoteId));
    final signed = bundle['signed_prekey'] as Map<String, dynamic>;
    final ot = bundle['one_time_prekey'] as Map<String, dynamic>?;
    final preKeyBundle = PreKeyBundle(
      1,
      deviceId,
      ot != null ? (ot['id'] as num).toInt() : 0,
      ot != null
          ? Curve25519PublicKey.fromBase64(ot['public'] as String)
          : Curve25519PublicKey.fromBase64(''),
      (signed['id'] as num).toInt(),
      Curve25519PublicKey.fromBase64(signed['public'] as String),
      base64Decode(signed['signature'] as String),
      IdentityKey.fromSerialized(base64Decode(bundle['identity_key'] as String)),
    );
    await builder.processPreKeyBundle(preKeyBundle);
  }

  Future<({int type, Uint8List body})> encrypt(
      String remoteId, String plaintext) async {
    await ensureIdentity();
    final cipher = SessionCipher(
        _store!, _store!, _store!, _store!, _addr(remoteId));
    final ct = await cipher.encrypt(utf8.encode(plaintext));
    return (type: ct.type, body: ct.serialize());
  }

  Future<String> decrypt(String remoteId, int type, Uint8List body) async {
    await ensureIdentity();
    final cipher = SessionCipher(
        _store!, _store!, _store!, _store!, _addr(remoteId));
    CiphertextMessage msg;
    if (type == 4) {
      msg = PreKeySignalMessage.fromSerialized(body);
      final plain = await cipher.decrypt(msg as PreKeySignalMessage);
      return utf8.decode(plain);
    }
    msg = SignalMessage.fromSerialized(body);
    final plain = await cipher.decrypt(msg as SignalMessage);
    return utf8.decode(plain);
  }
}
```

### `client/lib/net/api_client.dart`

```dart
import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart' as http;

import '../config/app_config.dart';

class ApiException implements Exception {
  final int status;
  final String code;
  final String message;
  ApiException(this.status, this.code, this.message);
  @override
  String toString() => '[$code] $message';
}

class ApiClient {
  final String base;
  String? token;
  final http.Client _client = http.Client();

  ApiClient({this.base = AppConfig.baseUrl, this.token});

  Map<String, String> get _headers => {
        'Content-Type': 'application/json',
        if (token != null) 'Authorization': 'Bearer $token',
      };

  Future<Map<String, dynamic>> _handle(http.Response resp) async {
    final body = resp.body.isEmpty
        ? <String, dynamic>{}
        : jsonDecode(resp.body) as Map<String, dynamic>;
    if (resp.statusCode >= 200 && resp.statusCode < 300) return body;
    final err = body['error'];
    throw ApiException(
      resp.statusCode,
      err is Map ? err['code']?.toString() ?? 'error' : 'error',
      err is Map ? err['message']?.toString() ?? '' : resp.body,
    );
  }

  Future<Map<String, dynamic>> getParams() async {
    final r = await _client.get(Uri.parse('$base/api/v1/params'));
    return _handle(r);
  }

  Future<Map<String, dynamic>> challenge(String username, String intent) async {
    final r = await _client.post(
      Uri.parse('$base/api/v1/auth/challenge'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'username': username, 'intent': intent}),
    );
    return _handle(r);
  }

  Future<Map<String, dynamic>> register(
      String username, String hash, String nonce, String proof) async {
    final r = await _client.post(
      Uri.parse('$base/api/v1/auth/register'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'username': username,
        'argon2id_hash': hash,
        'nonce': nonce,
        'proof': proof,
      }),
    );
    return _handle(r);
  }

  Future<Map<String, dynamic>> login(
      String username, String nonce, String proof) async {
    final r = await _client.post(
      Uri.parse('$base/api/v1/auth/login'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'username': username, 'nonce': nonce, 'proof': proof}),
    );
    return _handle(r);
  }

  Future<Map<String, dynamic>> uploadKeys(Map<String, dynamic> bundle) async {
    final r = await _client.post(Uri.parse('$base/api/v1/keys'),
        headers: _headers, body: jsonEncode(bundle));
    return _handle(r);
  }

  Future<Map<String, dynamic>> fetchBundle(String accountId) async {
    final r = await _client.get(Uri.parse('$base/api/v1/keys/$accountId'),
        headers: _headers);
    return _handle(r);
  }

  Future<Map<String, dynamic>> sendEnvelope({
    required String to,
    required int cipherType,
    required String body,
    required String clientMsgId,
  }) async {
    final r = await _client.post(Uri.parse('$base/api/v1/envelopes'),
        headers: _headers,
        body: jsonEncode({
          'to': to,
          'cipher_type': cipherType,
          'body': body,
          'client_msg_id': clientMsgId,
        }));
    return _handle(r);
  }

  Future<List<dynamic>> pullEnvelopes() async {
    final r = await _client.get(Uri.parse('$base/api/v1/envelopes'),
        headers: _headers);
    final m = await _handle(r);
    return (m['envelopes'] as List? ?? const []);
  }

  Future<void> ack(List<int> ids) async {
    if (ids.isEmpty) return;
    final r = await _client.post(Uri.parse('$base/api/v1/envelopes/ack'),
        headers: _headers, body: jsonEncode({'ids': ids}));
    await _handle(r);
  }

  Future<String> wsTicket() async {
    final r = await _client.post(Uri.parse('$base/api/v1/ws-ticket'),
        headers: _headers, body: jsonEncode({}));
    final m = await _handle(r);
    return m['ticket'] as String;
  }
}
```

### `client/lib/net/realtime.dart`

```dart
import 'dart:convert';

import 'package:web_socket_channel/web_socket_channel.dart';

import '../config/app_config.dart';

/// 单连接实时通道：收到 envelope 帧后回调；发送 ack 即完成投递闭环。
class Realtime {
  WebSocketChannel? _ch;
  bool _closed = false;

  void connect(String ticket, void Function(Map<String, dynamic> frame) onFrame) {
    final uri = Uri.parse('${AppConfig.wsUrl}?ticket=$ticket');
    final ch = WebSocketChannel.connect(uri);
    _ch = ch;
    ch.stream.listen((raw) {
      if (raw is String) {
        try {
          onFrame(jsonDecode(raw) as Map<String, dynamic>);
        } catch (_) {}
      }
    }, onError: (_) => close());
  }

  void ack(List<int> ids) {
    if (ids.isEmpty) return;
    _ch?.sink.add(jsonEncode({'type': 'ack', 'ids': ids}));
  }

  void close() {
    if (_closed) return;
    _closed = true;
    _ch?.sink.close();
  }
}
```

### `client/lib/state/chat_store.dart`

```dart
import 'dart:convert';

import 'package:flutter/foundation.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:uuid/uuid.dart';

import '../crypto/password.dart';
import '../crypto/signal_service.dart';
import '../net/api_client.dart';
import '../net/realtime.dart';

enum MsgStatus { sending, sent, delivered, read }

class ChatMsg {
  final String id;
  final bool fromMe;
  final String text;
  final MsgStatus status;
  final DateTime ts;
  const ChatMsg({
    required this.id,
    required this.fromMe,
    required this.text,
    required this.status,
    required this.ts,
  });
  ChatMsg copy({MsgStatus? status}) => ChatMsg(
        id: id, fromMe: fromMe, text: text,
        status: status ?? this.status, ts: ts,
      );
}

class ChatStore extends ChangeNotifier {
  final api = ApiClient();
  final signal = SignalService();
  final realtime = Realtime();
  final storage = const FlutterSecureStorage();

  String? accountId;
  String? contactId;
  final List<ChatMsg> messages = [];

  bool get loggedIn => accountId != null;

  Future<void> bootstrap() async {
    accountId = await storage.read(key: 'account_id');
    api.token = await storage.read(key: 'token');
    if (accountId != null && api.token != null) {
      try {
        await _openSocket();
      } catch (e) {
        if (kDebugMode) print('resume socket failed: $e');
      }
    }
    notifyListeners();
  }

  Future<void> authenticate(
      String username, String password, bool register) async {
    final params = await api.getParams();
    final pepper = params['pepper'] as String;
    final derived = AccountPassword.accountId(username, pepper);
    final nonceResp = await api.challenge(username, register ? 'register' : 'login');
    final nonce = nonceResp['nonce'] as String;
    final argon = Argon2Params.fromServer(params);
    // Argon2id 在 Isolate 里跑，不卡 UI。
    final hash = await AccountPassword.hashPassword(password, derived, argon);
    final proof = AccountPassword.proof(hash, nonce);

    final Map<String, dynamic> resp;
    if (register) {
      resp = await api.register(username, hash, nonce, proof);
      await signal.ensureIdentity();
      await api.uploadKeys(await signal.buildBundle());
    } else {
      resp = await api.login(username, nonce, proof);
    }
    accountId = resp['account_id'] as String;
    api.token = resp['token'] as String;
    await storage.write(key: 'account_id', value: accountId!);
    await storage.write(key: 'token', value: api.token!);
    notifyListeners();
    await _openSocket();
  }

  Future<void> _openSocket() async {
    final ticket = await api.wsTicket();
    realtime.connect(ticket, _onFrame);
  }

  Future<void> startChatWith(String targetId) async {
    contactId = targetId.trim().toLowerCase();
    final bundle = await api.fetchBundle(contactId!);
    await signal.establishSession(contactId!, bundle);
    final pend = await api.pullEnvelopes();
    final ackIds = <int>[];
    for (final e in pend) {
      final m = e as Map<String, dynamic>;
      await _consume(m);
      ackIds.add(m['id'] as int);
    }
    await api.ack(ackIds);
    notifyListeners();
  }

  Future<void> _consume(Map<String, dynamic> frame) async {
    final body = base64Decode(frame['body'] as String);
    final from = frame['from'] as String? ?? '';
    final plain = await signal.decrypt(from, frame['cipher_type'] as int, body);
    messages.add(ChatMsg(
      id: frame['client_msg_id'] as String? ?? const Uuid().v4(),
      fromMe: false,
      text: plain,
      status: MsgStatus.delivered,
      ts: DateTime.now(),
    ));
  }

  void _onFrame(Map<String, dynamic> frame) async {
    if (frame['type'] != 'envelope') return;
    await _consume(frame);
    realtime.ack([frame['id'] as int]);
    notifyListeners();
  }

  Future<void> sendText(String text) async {
    final target = contactId;
    if (target == null || text.trim().isEmpty) return;
    final id = const Uuid().v4();
    messages.add(ChatMsg(
        id: id, fromMe: true, text: text, status: MsgStatus.sending, ts: DateTime.now()));
    notifyListeners();
    try {
      final ct = await signal.encrypt(target, text);
      final r = await api.sendEnvelope(
        to: target,
        cipherType: ct.type,
        body: base64Encode(ct.body),
        clientMsgId: id,
      );
      final i = messages.indexWhere((m) => m.id == id);
      if (i >= 0) messages[i] = messages[i].copy(status: MsgStatus.sent);
      if (r['delivered'] == true) {
        messages[i] = messages[i].copy(status: MsgStatus.delivered);
      }
    } catch (e) {
      final i = messages.indexWhere((m) => m.id == id);
      if (i >= 0) messages[i] = messages[i].copy(status: MsgStatus.read);
    }
    notifyListeners();
  }

  Future<void> logout() async {
    realtime.close();
    await storage.delete(key: 'token');
    await storage.delete(key: 'account_id');
    accountId = null;
    messages.clear();
    notifyListeners();
  }
}
```

### `client/lib/ui/auth_page.dart`

```dart
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../state/chat_store.dart';

class AuthPage extends StatefulWidget {
  const AuthPage({super.key});

  @override
  State<AuthPage> createState() => _AuthPageState();
}

class _AuthPageState extends State<AuthPage> {
  final _user = TextEditingController();
  final _pass = TextEditingController();
  bool _register = true;
  bool _busy = false;
  String? _error;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: SingleChildScrollView(
          padding: const EdgeInsets.all(24),
          child: ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 360),
            child: Column(
              children: [
                Image.asset('assets/app_icon.png', width: 88, height: 88),
                const SizedBox(height: 16),
                Text('Paperplane', style: Theme.of(context).textTheme.headlineMedium),
                const Text('零知识 E2EE 聊天', style: TextStyle(color: Colors.grey)),
                const SizedBox(height: 28),
                TextField(
                  controller: _user,
                  decoration: const InputDecoration(labelText: '用户名'),
                  autocorrect: false,
                ),
                const SizedBox(height: 12),
                TextField(
                  controller: _pass,
                  decoration: const InputDecoration(labelText: '密码'),
                  obscureText: true,
                ),
                const SizedBox(height: 16),
                if (_error != null)
                  Text(_error!, style: const TextStyle(color: Colors.red)),
                const SizedBox(height: 8),
                FilledButton(
                  onPressed: _busy ? null : _submit,
                  child: Text(_busy ? '请稍候…' : (_register ? '注册并建立密钥会话' : '登录')),
                ),
                TextButton(
                  onPressed: () => setState(() => _register = !_register),
                  child: Text(_register ? '已有账号？去登录' : '没有账号？去注册'),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Future<void> _submit() async {
    setState(() {
      _busy = true;
      _error = null;
    });
    try {
      await context
          .read<ChatStore>()
          .authenticate(_user.text.trim(), _pass.text, _register);
    } catch (e) {
      setState(() => _error = '$e');
    } finally {
      if (mounted) setState(() => _busy = false);
    }
  }
}
```

### `client/lib/ui/chat_page.dart`

```dart
import 'package:flutter/material.dart';
import 'package:flutter_chat_ui/flutter_chat_ui.dart';
import 'package:provider/provider.dart';

import '../state/chat_store.dart';
import 'widgets/composer.dart';
import 'widgets/message_bubble.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Paperplane'),
        actions: [
          IconButton(
            onPressed: () => context.read<ChatStore>().logout(),
            icon: const Icon(Icons.logout),
          ),
        ],
      ),
      body: Consumer<ChatStore>(
        builder: (_, store, __) {
          if (store.contactId == null) {
            return Center(
              child: OutlinedButton.icon(
                onPressed: () => _openTargetDialog(context, store),
                icon: const Icon(Icons.person_add),
                label: const Text('输入对方账号 ID 开始聊天'),
              ),
            );
          }
          // ChatList：Flyer Chat 提供动画分页与底部输入栏容器。
          return ChatList(
            items: store.messages,
            itemBuilder: (item, index) =>
                MessageBubble(msg: item as ChatMsg),
            bottomWidget: Composer(onSend: store.sendText),
            isLastPage: true,
            typingIndicatorOptions: const TypingIndicatorOptions(),
          );
        },
      ),
    );
  }

  Future<void> _openTargetDialog(BuildContext context, ChatStore store) async {
    final ctrl = TextEditingController();
    final id = await showDialog<String>(
      context: context,
      builder: (_) => AlertDialog(
        title: const Text('目标账号 ID'),
        content: TextField(
          controller: ctrl,
          decoration: const InputDecoration(hintText: '64 位 hex account id'),
        ),
        actions: [
          TextButton(
              onPressed: () => Navigator.pop(context), child: const Text('取消')),
          FilledButton(
              onPressed: () => Navigator.pop(context, ctrl.text),
              child: const Text('开始')),
        ],
      ),
    );
    if (id == null || id.isEmpty) return;
    try {
      await store.startChatWith(id);
    } catch (e) {
      if (context.mounted) {
        ScaffoldMessenger.of(context)
            .showSnackBar(SnackBar(content: Text('会话建立失败: $e')));
      }
    }
  }
}
```

### `client/lib/ui/theme.dart`

```dart
import 'package:flutter/material.dart';

// Material 3 双主题，蓝灰主色，深色/浅色跟随系统。
ThemeData buildLightTheme() => ThemeData(
      useMaterial3: true,
      colorScheme: ColorScheme.fromSeed(
        seedColor: Colors.blueGrey,
        brightness: Brightness.light,
      ),
      scaffoldBackgroundColor: const Color(0xFFF4F6FA),
      filledButtonTheme: FilledButtonThemeData(
        style: FilledButton.styleFrom(
          minimumSize: const Size.fromHeight(48),
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(14)),
        ),
      ),
      inputDecorationTheme: InputDecorationTheme(
        filled: true,
        fillColor: Colors.white,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(14),
          borderSide: const BorderSide(color: Color(0xFFE2E6EE)),
        ),
      ),
    );

ThemeData buildDarkTheme() => ThemeData(
      useMaterial3: true,
      colorScheme: ColorScheme.fromSeed(
        seedColor: Colors.blueGrey,
        brightness: Brightness.dark,
      ),
      scaffoldBackgroundColor: const Color(0xFF10141A),
      filledButtonTheme: FilledButtonThemeData(
        style: FilledButton.styleFrom(
          minimumSize: const Size.fromHeight(48),
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(14)),
        ),
      ),
    );

// 聊天气泡渐变（发送）/ 柔和浅色（接收）
const bubbleGradient = LinearGradient(
  colors: [Color(0xFF4E6E8E), Color(0xFF6C8EB0)],
  begin: Alignment.topLeft,
  end: Alignment.bottomRight,
);
const bubbleRadius = 18.0;
```

### `client/lib/ui/widgets/composer.dart`

```dart
import 'package:flutter/material.dart';

class Composer extends StatefulWidget {
  final void Function(String) onSend;
  const Composer({super.key, required this.onSend});

  @override
  State<Composer> createState() => _ComposerState();
}

class _ComposerState extends State<Composer> {
  final _ctrl = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Container(
        padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
        color: Theme.of(context).scaffoldBackgroundColor,
        child: Row(
          children: [
            IconButton(
              onPressed: () {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('图片 / 文件即将支持，当前为文本消息')),
                );
              },
              icon: const Icon(Icons.attach_file),
            ),
            Expanded(
              child: TextField(
                controller: _ctrl,
                minLines: 1,
                maxLines: 5,
                textInputAction: TextInputAction.send,
                decoration: const InputDecoration(
                  hintText: '加密消息',
                  border: OutlineInputBorder(),
                ),
                onSubmitted: (v) => _submit(),
              ),
            ),
            const SizedBox(width: 8),
            IconButton.filled(
              onPressed: _submit,
              icon: const Icon(Icons.send),
            ),
          ],
        ),
      ),
    );
  }

  void _submit() {
    final t = _ctrl.text.trim();
    if (t.isEmpty) return;
    widget.onSend(t);
    _ctrl.clear();
  }
}
```

### `client/lib/ui/widgets/message_bubble.dart`

```dart
import 'package:flutter/material.dart';

import '../state/chat_store.dart';
import 'theme.dart';

class MessageBubble extends StatelessWidget {
  final ChatMsg msg;
  const MessageBubble({super.key, required this.msg});

  @override
  Widget build(BuildContext context) {
    final isMe = msg.fromMe;
    return Align(
      alignment: isMe ? Alignment.centerRight : Alignment.centerLeft,
      child: Container(
        margin: const EdgeInsets.symmetric(vertical: 3, horizontal: 12),
        padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 14),
        constraints: BoxConstraints(
          maxWidth: MediaQuery.of(context).size.width * 0.75,
        ),
        decoration: BoxDecoration(
          gradient: isMe ? bubbleGradient : null,
          color: isMe ? null : Colors.white,
          borderRadius: BorderRadius.circular(bubbleRadius),
          boxShadow: const [
            BoxShadow(blurRadius: 4, color: Colors.black12, offset: Offset(0, 1)),
          ],
        ),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            Flexible(
              child: Text(
                msg.text,
                style: TextStyle(
                  color: isMe ? Colors.white : const Color(0xFF1B2430),
                  fontSize: 15,
                ),
              ),
            ),
            const SizedBox(width: 6),
            if (isMe) Icon(_tick(msg.status), size: 14, color: Colors.white70),
          ],
        ),
      ),
    );
  }

  IconData _tick(MsgStatus s) => switch (s) {
        MsgStatus.sending => Icons.access_time,
        MsgStatus.sent => Icons.check,
        MsgStatus.delivered => Icons.done_all,
        MsgStatus.read => Icons.done_all,
      };
}
```

### `client/web/index.html`

```html
<!DOCTYPE html>
<html>
<head>
  <base href="/" />
  <meta charset="UTF-8" />
  <meta content="IE=Edge" http-equiv="X-UA-Compatible" />
  <meta name="description" content="Paperplane — zero-knowledge E2EE chat" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0" />
  <title>Paperplane</title>
  <link rel="icon" type="image/png" href="assets/app_icon.png" />
</head>
<body>
  <script src="main.dart.js" defer></script>
</body>
</html>
```

### `config/config.dev.yaml`

```yaml
# 开发用配置：仅本地 docker-compose.dev.yml 用，占位密钥允许。
project: paperplane
environment: development

server:
  host: 127.0.0.1
  port: 8080
  public_url: http://127.0.0.1:8080
  ws_allowed_origins:
    - http://127.0.0.1:8080
    - http://localhost:8080
    - chrome-extension://*

database:
  host: 127.0.0.1
  port: 5432
  name: paperplane
  user: paperplane
  password: dev-password
  max_conns: 8
  envelope_ttl_hours: 24

redis:
  host: 127.0.0.1
  port: 6379
  password: ""
  db: 0

jwt:
  secret: dev-only-secret-at-least-32-bytes-long-xxxx
  ttl: 24h

crypto:
  pepper: dev-only-pepper
  argon2:
    memory_kib: 65536
    iterations: 3
    parallelism: 1

client:
  dev_server_url: http://127.0.0.1:8080
  app_name: Paperplane

defense:
  xdp:
    enabled: false
    iface: wlan0
  rate_limits:
    ws_per_ip: 32
    prekey_per_min: 120
    envelope_per_min: 240
    envelope_distinct_targets: 64
    envelope_window: 60s
    challenge_per_min: 30
  kotlin:
    redis_stream: paperplane:events
    consumer_group: paperplane-policy
    window: 60s
  crowdsec:
    collections: [crowdsecurity/http-cve]
    lapi_port: 8081
```

### `config/config.yaml`

```yaml
# =============================================================================
# Paperplane — 唯一配置源 (single source of truth)
# -----------------------------------------------------------------------------
# 这个文件是整个系统唯一的服务器配置真相源。
# 修改后执行 `make config`（即 scripts/render-config.sh），会派生：
#   deploy/.env                 -> docker compose 变量
#   deploy/caddy.env            -> Caddyfile 变量
#   deploy/crowdsec/acquis.yaml -> CrowdSec 日志采集
#   policy-sidecar/config/application.yaml -> Kotlin sidecar
#   xdp/config.env              -> XDP loader 参数
# Go 服务启动时直接读取本文件并 fail-fast 校验必填项。
# 任何服务都不允许再单独维护一份配置副本。
# =============================================================================

project: paperplane
environment: development   # production | development。生产部署前必须改成 production 并替换所有 change-me 占位

server:
  host: 127.0.0.1         # 只绑回环：公网流量必须经 Caddy 进入
  port: 8080
  public_url: https://chat.example.com
  ws_allowed_origins:
    - https://chat.example.com
  read_header_timeout: 10s
  idle_timeout: 120s
  shutdown_timeout: 15s

database:
  host: 127.0.0.1
  port: 5432
  name: paperplane
  user: paperplane
  password: change-me-postgres
  max_conns: 16
  envelope_ttl_hours: 24   # 离线密文最长保留时间，投递/确认后立即删除

redis:
  host: 127.0.0.1
  port: 6379
  password: change-me-redis
  db: 0
  pool_size: 16

jwt:
  secret: change-me-please-32-bytes-minimum-secret
  ttl: 24h
  issuer: paperplane

crypto:
  pepper: change-me-project-pepper-at-least-16-bytes
  argon2:
    memory_kib: 65536      # 64 MiB
    iterations: 3
    parallelism: 1

client:
  dev_server_url: http://127.0.0.1:8080
  app_name: Paperplane

defense:
  xdp:
    enabled: true
    iface: eth0
  rate_limits:
    ws_per_ip: 6           # 单 IP 并发 WebSocket 连接上限
    prekey_per_min: 30     # 预密钥包请求 / 分钟 / 账号
    envelope_per_min: 60   # 密文信封 / 分钟 / 账号
    envelope_distinct_targets: 20  # 窗口内不同收件人上限
    envelope_window: 60s
    challenge_per_min: 10  # 登录/注册挑战 / 分钟 / IP
  kotlin:
    redis_stream: paperplane:events
    consumer_group: paperplane-policy
    window: 60s
  crowdsec:
    collections:
      - crowdsecurity/http-cve
      - crowdsecurity/base-http-scenarios
    lapi_port: 8081          # 避开服务端 8080
```

### `deploy/.env`

```text
# 由 scripts/render-config.sh 从 config/config.yaml 生成，请勿手改。
COMPOSE_PROJECT_NAME=paperplane
PUBLIC_URL=https://chat.example.com
SERVER_HOST=127.0.0.1
SERVER_PORT=8080
UPSTREAM=127.0.0.1:8080
POSTGRES_DB=paperplane
POSTGRES_USER=paperplane
POSTGRES_PASSWORD=change-me-postgres
POSTGRES_HOST=127.0.0.1
POSTGRES_PORT=5432
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=change-me-redis
WS_ALLOWED_ORIGINS=https://chat.example.com
CONFIG_PATH=/etc/paperplane/config.yaml
CROWDSEC_LAPI_PORT=8081
KOTLIN_REDIS_STREAM=paperplane:events
KOTLIN_CONSUMER_GROUP=paperplane-policy
KOTLIN_WINDOW_SECONDS=60
WS_PER_IP=6
PREKEY_PER_MIN=30
ENVELOPE_PER_MIN=60
ENVELOPE_DISTINCT_TARGETS=20
ENVELOPE_WINDOW_SECONDS=60
CHALLENGE_PER_MIN=10
```

### `deploy/Caddyfile`

```text
{$PUBLIC_URL} {
	encode gzip

	header {
		Strict-Transport-Security "max-age=63072000; includeSubDomains"
		X-Content-Type-Options "nosniff"
		X-Frame-Options "DENY"
		Referrer-Policy "no-referrer"
		Permissions-Policy "interest-cohort=()"
		-Server
	}

	log {
		output file /var/log/caddy/access.log {
			roll_size 10MiB
			roll_keep 3
			roll_keep_for_days 7
		}
		format json
	}

	# WebSocket 原生支持：Upgrade 自动透传，Go 服务端只校验 Origin 白名单
	reverse_proxy {$UPSTREAM} {
		header_up X-Forwarded-For {remote_host}
		header_up X-Forwarded-Proto {scheme}
		flush_interval -1
		transport http {
			dial_timeout 3s
			response_header_timeout 30s
		}
	}
}
```

### `deploy/caddy.env`

```dotenv
PUBLIC_URL=https://chat.example.com
UPSTREAM=127.0.0.1:8080
```

### `deploy/docker-compose.yml`

```yaml
# Paperplane 生产编排 — 变量全部来自 render-config.sh 生成的 .env / caddy.env
# 架构：所有服务 network_mode: host（LInux 单机部署），服务端绑 127.0.0.1:8080，
# Caddy 在本机 80/443 终止 TLS。Postgres/Redis 仅暴露 127.0.0.1。
services:
  postgres:
    image: postgres:16-alpine
    restart: unless-stopped
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    ports:
      - "127.0.0.1:5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB}"]
      interval: 5s
      timeout: 5s
      retries: 20

  redis:
    image: redis:7-alpine
    restart: unless-stopped
    command: ["redis-server", "--requirepass", "${REDIS_PASSWORD}", "--port", "6379", "--maxmemory", "256mb", "--maxmemory-policy", "allkeys-lru"]
    ports:
      - "127.0.0.1:6379:6379"
    volumes:
      - redisdata:/data
    healthcheck:
      test: ["CMD", "redis-cli", "-a", "${REDIS_PASSWORD}", "ping"]
      interval: 5s
      timeout: 5s
      retries: 20

  server:
    build: ../server
    restart: unless-stopped
    network_mode: host
    depends_on:
      postgres: { condition: service_healthy }
      redis: { condition: service_healthy }
    volumes:
      - ../config/config.yaml:${CONFIG_PATH:-/etc/paperplane/config.yaml}:ro
      - /var/lib/paperplane:/var/lib/paperplane
    environment:
      CONFIG_PATH: ${CONFIG_PATH:-/etc/paperplane/config.yaml}
      MIGRATIONS_DIR: /migrations
      BAN_FILE: /var/lib/paperplane/bans.txt

  caddy:
    image: caddy:2-alpine
    restart: unless-stopped
    network_mode: host
    depends_on: [server]
    env_file:
      - caddy.env
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - caddydata:/data
      - ./caddylogs:/var/log/caddy

  crowdsec:
    image: crowdsecurity/crowdsec:latest
    restart: unless-stopped
    network_mode: host
    depends_on: [caddy]
    volumes:
      - ./crowdsec/config.yaml.local:/etc/crowdsec/config.yaml.local:ro
      - ./crowdsec/acquis.yaml:/etc/crowdsec/acquis.yaml:ro
      - crowdsec-data:/var/lib/crowdsec/data
      - ./caddylogs:/var/log/caddy:ro
    command:
      - sh
      - -c
      - "cscli collections install crowdsecurity/http-cve crowdsecurity/base-http-scenarios --force 2>/dev/null || true; exec crowdsec -c /etc/crowdsec/config.yaml"

  kotlin-policy:
    build: ../policy-sidecar
    restart: unless-stopped
    network_mode: host
    depends_on:
      redis: { condition: service_healthy }
    volumes:
      - ../policy-sidecar/config/application.yaml:/app/config/application.yaml:ro
    environment:
      CROWDSEC_LAPI_KEY: ${CROWDSEC_LAPI_KEY:-}

volumes:
  pgdata:
  redisdata:
  caddydata:
  crowdsec-data:
```

### `deploy/crowdsec/acquis.yaml`

```yaml
# 由 render-config.sh 生成：Caddy 访问日志 -> CrowdSec 场景引擎
filenames:
  - /var/log/caddy/access.log
labels:
  type: caddy
```

### `deploy/crowdsec/config.yaml.local`

```text
# CrowdSec 本地 API 配置（由 render-config.sh 生成时保留；端口 8081 避开 8080）
api:
  server:
    listen: 127.0.0.1:8081
  client:
    credentials:
      user: ci
      password: change-me-cs-password

# 关键规则：WebSocket 长连接不允许被 bouncer 误杀
common:
  max_upload_sources: 5
  online_client:
    enabled: false

# 阻止 bouncer 对 WebSocket 长连接下发 ban：以 source 白名单形式排除
# （Caddy 的 X-Forwarded-For 与服务端 IP 限流同时生效）
```

### `docs/ARCHITECTURE.md`

```markdown
# 架构 (ARCHITECTURE.md)

## 组件

| 组件 | 职责 |
|---|---|
| Flutter client | 账号派生 / libsignal 会话 / UI / 本地加密存储 |
| Go server | HTTP + WS 网关、零知识信封转发、内嵌限流 |
| PostgreSQL | 账号表、预密钥包、待投信封 |
| Redis | 一次性 nonce / 票据、限流计数、事件流、封禁 ZSET |
| Kotlin sidecar | 消费 Redis Streams，跨维规则，下发封禁 |
| Caddy | TLS 终止 + WS 反代 + 访问日志 |
| CrowdSec | 场景引擎 + LAPI 决策 |
| XDP loader | 把封禁 ZSET 同步到内核 LPM trie |

## 请求流

```
Client ──HTTPS──▶ Caddy ──▶ Go:8080 ──▶ Postgres / Redis
                       ▲
Client ──WSS───────────┘
                │
                ├─ Go 发布事件 ──▶ Redis Stream
                │                       │
                │                  Kotlin sidecar
                │                       │ ban:ip ZSET
                ▼                       ▼
            CrowdSec LAPI          XDP loader → LPM map
```

## 数据流

- 注册：`username+password → PHC → 存(account_id, PHC)`
- 发消息：`明文 → Signal.encrypt → 信封 → 转发 → ack 后删除`
- 在线推送：`Hub.Deliver → WebSocket frame → 客户端 ack`

## 配置流

```
config/config.yaml
        │
        ▼ scripts/render-config.sh
   ├── deploy/.env
   ├── deploy/caddy.env
   ├── deploy/crowdsec/acquis.yaml
   ├── policy-sidecar/config/application.yaml
   └── xdp/config.env
```

任何服务不允许再自己维护第二份配置。
```

### `docs/CRYPTO.md`

```markdown
# 密码学模型 (CRYPTO.md)

## 1. 账号单向派生

```
account_id = SHA-256( username.lower() ‖ pepper )           # hex64
argon2id   = Argon2id( password, salt = account_id[0:16],
                       m = 64 MiB, t = 3, p = 1 )           # PHC 字符串
```

- `pepper` 通过 `GET /api/v1/params` 公开下发，只做域名分隔，不是秘密。
- 服务端只存 `account_id` + `argon2id_hash`，永不接触明文密码。
- 登录：客户端拿一次性 `nonce`，证明 `HMAC-SHA256(credential, nonce)`。
  抓包者拿到的 JWT / nonce 都不能重放。

## 2. E2EE（X3DH + Double Ratchet）

- 首次登录生成：
  - IdentityKey (Curve25519 密钥对)
  - SignedPreKey (签名公钥)
  - 100 个 OneTimePreKey
  - KyberPreKey（预留字段，当前 libsignal 纯 Dart 版尚未启用 PQXDH）
- 公钥上传，私钥仅存 `flutter_secure_storage`。
- 会话建立：`SessionBuilder.processPreKeyBundle(remoteBundle)`。
- 收发：`SessionCipher.encrypt / decrypt`；密文只以 `base64` 信封过服务端。

## 3. 信封协议

```
POST /api/v1/envelopes
{
  "to": "<account_id hex64>",
  "cipher_type": 3 | 4,          // 3=Whisper, 4=PreKeyWhisper
  "body": "<base64(SignalMessage)>",
  "client_msg_id": "<uuid>",
  "sealed_sender": true|false,
  "sender_token": "<盲令牌>"
}
```

- 服务端不解析 `body`，只按 `to` 转发。
- 投递成功即 `sender_id = NULL`（清除路由元数据）。
- 收件人 `ack` 后整行删除；离线密文 24h TTL 兜底清理。

## 4. Sealed Sender

- `POST /api/v1/sealed-tokens` 一次领 5 个盲令牌（24h 有效）。
- 信封携带 `sender_token`；服务端验票后**不记录**发送方账号。
- 收件人只看到密文，不知道是谁发的。

## 5. 零知识承诺

| 服务端允许 | 服务端禁止 |
|---|---|
| 转发密文信封 | 解密/记录密文 |
| 存 account_id + PHC | 存可逆用户名 |
| 在线推送 + ack 删除 | 内存 Map 当生产 store |
| 单点限流与事件统计 | 记录 JWT / Origin / 口令 |

## 6. 不要做的事

- 不要自研 E2EE；只允许集成 libsignal。
- 不要把 Argon2 跑在 UI 线程；必须 `Isolate.run`。
- 不要在客户端存 `pepper` 以外的密钥；不要持久化会话状态。
- 不要在 WS / HTTP 处理里 `return true` 跳过 Origin 校验。
```

### `docs/KNOWN_ERRORS.md`

```markdown
# 已知错误与防御清单 (KNOWN_ERRORS.md)

## A – J：协议 / 服务端层

- **A. 会话状态必须事务提交后才返回**：bundle 取 OTP 必须在事务里 `SELECT ... FOR UPDATE SKIP LOCKED` 后再 `DELETE`，commit 完才给客户端 200。
- **B. 按 ProtocolAddress 加 `synchronized`**：对 `SignalProtocolStore` 实例级串行化，避免两个设备并发写同一会话状态。
- **C. 跳过密钥依赖 libsignal 内置**：不要自己跳过 ratchet；密钥交换失败让协议自己报 `DuplicateMessageException`。
- **D. MAC 失败提供重置会话**：`decrypt` 抛 `InvalidMessageException` 时删本地会话、重取 bundle 重发。
- **E. 跨平台序列化 schema 一致**：密文信封字段名 `cipher_type` / `body` / `client_msg_id` 两端对齐；客户端 base64 标准编码。
- **F. WS Origin 精确白名单**：空 Origin 放行（原生客户端），非空必须精确匹配 `scheme://host:port`，不要用 `strings.HasPrefix`。
- **G. Caddy 后 WS 不通**：不要用 `r.Host` 比 Origin；Caddy 会改 Host 但保留 `Origin`，只信 `Origin` 头。
- **H. 服务端绑回环**：只监听 `127.0.0.1:8080`，公网入口必须经 Caddy。
- **I. Argon2id 入 `Isolate.run`**：UI 线程不能跑 m=64MiB；慢哈希必须后台。
- **J. 自托管 Runner 构建 403**：CI 用 `GOTOOLCHAIN=local`，别在公开 Runner 上拉私有工具链。

## K-1 – K-9：内核 / 策略层

- **K-1. XDP virtio 回退 generic**：virtio 网卡不支持 native XDP，`ethtool -i` 看 driver；先尝试 `XDP_FLAGS_SKB_MODE`，失败再 `DRV_MODE`。
- **K-2. BPF map flock 串行化**：loader 用单线程读写 `ban_map`；多线程同时 `update_elem` 会破坏 LPM trie 内部结构。
- **K-3. XDP detach 显式卸载**：`SIGTERM/SIGINT/SIGHUP` 必须 `bpf_xdp_detach(ifindex, flags, NULL)`，不要依赖进程退出自动清理。
- **K-4. C++ ASan/UBSan**：CI 里 `clang++ -fsanitize=address,undefined` 跑 loader，禁止裸 `new[]/delete[]` 错配。
- **K-5. Kotlin 协程结构化并发**：顶层用 `coroutineScope` + `SupervisorJob`，单条消息异常不能杀消费循环。
- **K-6. Redis 故障 fail-open 50ms 超时**：限流/封禁查询一律 `context.WithTimeout(50ms)`，Redis 挂时服务照常（降级为无防护）。
- **K-7. 时钟统一 Redis TIME**：sidecar 用 `redis.time()` 做时间戳，不要本机时钟漂移导致规则误判。
- **K-8. bouncer 缓存+长连接**：CrowdSec bouncer 用 `cache_duration` 缓存决策，避免每个连接打 LAPI。
- **K-9. never-ban 排除长连接**：WebSocket 长连接在 CrowdSec 里走白名单（`source` 排除），不要让 bouncer 因 idle 误封正常客户端。
```

### `docs/SELF_HOSTING.md`

```markdown
# 自托管指南 (SELF_HOSTING.md)

## 0. 前置

- 一台 Linux 服务器（推荐 2C2G），域名已解析到本机。
- Docker 24+ / docker compose 插件。
- 开放 80 / 443（Caddy 自动签 Let's Encrypt）。

## 1. 拉代码

```bash
git clone https://github.com/<your-org>/paperplane-e2ee.git
cd paperplane-e2ee
```

## 2. 改唯一配置

编辑 `config/config.yaml`：

- `server.public_url`: `https://chat.example.com`
- `server.ws_allowed_origins`: 你的域名
- `jwt.secret`: `openssl rand -hex 32`
- `crypto.pepper`: `openssl rand -hex 16`
- `database.password` / `redis.password`: 随机
- `defense.xdp.iface`: 公网网卡名（`ip -o link`）

```bash
make config-check        # 校验必填
make config              # 生成 deploy/.env / caddy.env / acquis.yaml / sidecar 配置
```

## 3. 启动

```bash
cd deploy
CROWDSEC_LAPI_KEY=$(openssl rand -hex 24) docker compose up -d
```

服务端监听 `127.0.0.1:8080`；Caddy 占 80/443；CrowdSec LAPI 在 `127.0.0.1:8081`。

## 4. 验证

```bash
curl https://chat.example.com/healthz
# {"status":"ok","uptime_s":3,"time":"..."}
```

客户端里：把 `SERVER_URL` 指向 `https://chat.example.com` 即可。

## 5. XDP 内核过滤（可选）

```bash
sudo apt install -y clang libbpf-dev linux-tools-common
cd xdp && make
sudo ./xdp_loader eth0 /var/lib/paperplane/bans.txt
```

封禁 IP 会由 Go 服务端定期写入 `bans.txt`，loader 每 5s 同步到 LPM trie。

## 6. 备份与升级

- 只需要备份 Postgres：`pg_dump paperplane > backup.sql`。
- Redis 不存任何持久数据，重启即清。
- 升级：`git pull && make config && cd deploy && docker compose up -d --build`。

## 常见问题

- **WS 连不上**：先看 Caddy 日志 `/var/log/caddy/access.log`，确认 `Upgrade` 头被转发；不要把 Origin 白名单写错。
- **注册 409**：用户名已存在；要重新注册先换用户名（当前不提供账号恢复/删除接口）。
- **客户端连不上**：浏览器跨域需把 `https://chat.example.com` 加进 `ws_allowed_origins`，Caddy 必须 HTTPS。
```

### `docs/_config.yml`

```yaml
remote_theme: pages-themes/cayman@v0.2.0
plugins:
  - jekyll-remote-theme
title: Paperplane
description: 零知识 E2EE 加密聊天系统 · AGPL-3.0
show_downloads: false
google_analytics:
```

### `docs/index.html`

```html
<!doctype html>
<html lang="zh-CN">
<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<title>Paperplane — 零知识 E2EE 加密聊天</title>
<link rel="icon" type="image/png" href="assets/logo.png" />
<style>
  :root {
    --bg: #0b0f16;
    --panel: #121826;
    --line: rgba(255,255,255,0.08);
    --text: #e7ecf3;
    --muted: #9aa7b8;
    --accent: #6fa8dc;
    --accent2: #8fd3b8;
    --warn: #f0b868;
  }
  * { box-sizing: border-box; }
  html, body { margin: 0; padding: 0; }
  body {
    background: radial-gradient(1200px 600px at 80% -10%, #1b2740 0%, transparent 60%),
                radial-gradient(900px 500px at -10% 20%, #16222e 0%, transparent 60%),
                var(--bg);
    color: var(--text);
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC",
                 "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
    line-height: 1.6;
    -webkit-font-smoothing: antialiased;
  }
  .wrap { max-width: 1080px; margin: 0 auto; padding: 0 24px; }

  /* Nav */
  nav {
    position: sticky; top: 0; z-index: 10;
    backdrop-filter: blur(12px);
    background: rgba(11,15,22,0.7);
    border-bottom: 1px solid var(--line);
  }
  nav .wrap { display: flex; align-items: center; justify-content: space-between; height: 60px; }
  .brand { display: flex; align-items: center; gap: 10px; font-weight: 600; }
  .brand img { width: 28px; height: 28px; border-radius: 6px; }
  nav a { color: var(--muted); text-decoration: none; margin-left: 18px; font-size: 14px; }
  nav a:hover { color: var(--text); }

  /* Hero */
  header { padding: 90px 0 70px; }
  .hero { display: grid; grid-template-columns: 1.1fr 0.9fr; gap: 40px; align-items: center; }
  h1 { font-size: 44px; line-height: 1.15; margin: 0 0 16px; letter-spacing: -0.5px; }
  h1 .grad {
    background: linear-gradient(135deg, var(--accent), var(--accent2));
    -webkit-background-clip: text; background-clip: text; color: transparent;
  }
  .sub { color: var(--muted); font-size: 17px; margin: 0 0 28px; max-width: 560px; }
  .cta { display: flex; gap: 12px; flex-wrap: wrap; }
  .btn {
    display: inline-block; padding: 12px 22px; border-radius: 10px;
    font-size: 15px; text-decoration: none; border: 1px solid transparent;
    transition: transform .15s ease, background .15s ease;
  }
  .btn-primary { background: linear-gradient(135deg, var(--accent), #4a7bb5); color: #fff; }
  .btn-ghost { border-color: var(--line); color: var(--text); background: rgba(255,255,255,0.03); }
  .btn:hover { transform: translateY(-1px); }

  /* Dove illustration */
  .dove-card {
    background: linear-gradient(160deg, rgba(111,168,220,0.12), rgba(143,211,184,0.08));
    border: 1px solid var(--line); border-radius: 20px;
    padding: 30px; display: flex; align-items: center; justify-content: center;
  }
  .dove-card svg { width: 100%; max-width: 320px; height: auto; }

  /* Section */
  section { padding: 60px 0; border-top: 1px solid var(--line); }
  h2 { font-size: 28px; margin: 0 0 8px; }
  .lead { color: var(--muted); margin: 0 0 28px; }

  /* Cards */
  .grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 16px; }
  .card {
    background: var(--panel); border: 1px solid var(--line); border-radius: 14px;
    padding: 20px;
  }
  .card h3 { margin: 0 0 8px; font-size: 16px; }
  .card p { margin: 0; color: var(--muted); font-size: 14px; }
  .card .tag {
    display: inline-block; font-size: 12px; color: var(--accent2);
    background: rgba(143,211,184,0.1); padding: 2px 8px; border-radius: 999px; margin-bottom: 10px;
  }

  /* Model diagram */
  .model {
    background: var(--panel); border: 1px solid var(--line); border-radius: 16px;
    padding: 24px; overflow-x: auto;
  }
  .row { display: flex; gap: 12px; align-items: stretch; flex-wrap: wrap; }
  .node {
    flex: 1; min-width: 140px; text-align: center;
    padding: 14px 10px; border-radius: 10px; border: 1px solid var(--line);
    background: rgba(255,255,255,0.02);
  }
  .node b { display: block; font-size: 14px; margin-bottom: 4px; }
  .node span { color: var(--muted); font-size: 12px; }
  .arrow { align-self: center; color: var(--muted); font-size: 20px; }

  /* Code block */
  pre {
    background: #0a0e15; border: 1px solid var(--line); border-radius: 12px;
    padding: 16px; overflow-x: auto; font-size: 13px; line-height: 1.55;
    color: #c8d3e0;
  }
  code { font-family: ui-monospace, SFMono-Regular, Menlo, monospace; }

  /* Footer */
  footer { padding: 40px 0; border-top: 1px solid var(--line); color: var(--muted); font-size: 13px; }
  footer a { color: var(--accent); text-decoration: none; }

  @media (max-width: 820px) {
    .hero { grid-template-columns: 1fr; }
    h1 { font-size: 32px; }
    .grid { grid-template-columns: 1fr; }
    nav a { display: none; }
  }
</style>
</head>
<body>

<nav>
  <div class="wrap">
    <div class="brand">
      <img src="assets/logo.png" alt="logo" />
      <span>Paperplane</span>
    </div>
    <div>
      <a href="#features">特性</a>
      <a href="#model">安全模型</a>
      <a href="#config">配置</a>
      <a href="CRYPTO.md">密码学</a>
      <a href="SELF_HOSTING.md">自托管</a>
      <a href="KNOWN_ERRORS.md">错误清单</a>
      <a class="btn btn-primary" style="margin-left:18px;" href="https://github.com/paperplane-e2ee/paperplane-e2ee">GitHub</a>
    </div>
  </div>
</nav>

<header>
  <div class="wrap hero">
    <div>
      <h1>纸飞机风格的<br /><span class="grad">零知识 E2EE 加密聊天</span></h1>
      <p class="sub">
        服务端不解密、不记录密文、不存储可逆账号字段；私钥仅客户端安全存储。
        Flutter + Go + libsignal + Kotlin sidecar + eBPF/XDP，单一配置源，AGPL-3.0。
      </p>
      <div class="cta">
        <a class="btn btn-primary" href="https://github.com/paperplane-e2ee/paperplane-e2ee">在 GitHub 上查看</a>
        <a class="btn btn-ghost" href="SELF_HOSTING.md">五分钟自托管</a>
      </div>
    </div>
    <div class="dove-card" aria-hidden="true">
      <!-- 黑白鸽子图标（单色 SVG 示意，实际用 assets/logo.png） -->
      <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
        <path fill="#ffffff" d="M40 110c10-30 45-45 75-40 25 4 45 20 55 35-15 5-30 5-45 0-5 15-20 25-40 25-15 0-35-8-45-20z"/>
        <path fill="#0b0f16" d="M60 118c15-20 45-30 75-22-10 12-30 22-55 22-8 0-15 0-20 0z"/>
        <circle cx="150" cy="100" r="3" fill="#ffffff"/>
      </svg>
    </div>
  </div>
</header>

<section id="features">
  <div class="wrap">
    <h2>为隐私而生的分层防御</h2>
    <p class="lead">从应用协议到内核网络栈，每一层都假设对手已经拿到数据库。</p>
    <div class="grid">
      <div class="card">
        <span class="tag">E2EE</span>
        <h3>libsignal X3DH + Double Ratchet</h3>
        <p>不自研密码学。IdentityKey / SignedPreKey / OneTimePreKey 在客户端生成，私钥入 flutter_secure_storage。</p>
      </div>
      <div class="card">
        <span class="tag">Zero-Knowledge</span>
        <h3>只存账号 ID 与 PHC</h3>
        <p>account_id = SHA-256(username ‖ pepper)，登录用 HMAC(nonce) 挑战；服务端永不接触明文密码。</p>
      </div>
      <div class="card">
        <span class="tag">Sealed Sender</span>
        <h3>盲令牌发信</h3>
        <p>信封投递后清除路由元数据，ack 即删除；离线密文 24h TTL 兜底，服务端不知道谁发给谁。</p>
      </div>
      <div class="card">
        <span class="tag">Runtime Defense</span>
        <h3>多层限流</h3>
        <p>WS Origin 精确白名单、每 IP 连接数、预密钥与信封令牌桶、distinct-target 滑动窗口。</p>
      </div>
      <div class="card">
        <span class="tag">Policy Sidecar</span>
        <h3>Kotlin + Redis Streams</h3>
        <p>结构化并发消费事件，Redis TIME 时钟，跨维规则（IP→账号比值、重连频率）下发封禁。</p>
      </div>
      <div class="card">
        <span class="tag">Kernel</span>
        <h3>eBPF/XDP LPM 黑名单</h3>
        <p>封禁 IP 在内核态直接丢包，SIGTERM 显式 detach，RAII 封装 libbpf，CI 启用 ASan/UBSan。</p>
      </div>
    </div>
  </div>
</section>

<section id="model">
  <div class="wrap">
    <h2>安全模型一览</h2>
    <p class="lead">客户端永远是可信边界；服务端只做信封转发与准入控制。</p>
    <div class="model">
      <div class="row">
        <div class="node">
          <b>Flutter 客户端</b>
          <span>libsignal<br />secure storage</span>
        </div>
        <div class="arrow">⇄ E2EE ⇄</div>
        <div class="node">
          <b>Go Relay</b>
          <span>只看信封<br />不看明文</span>
        </div>
        <div class="arrow">⇄</div>
        <div class="node">
          <b>Postgres</b>
          <span>account_id + PHC<br />待投信封</span>
        </div>
      </div>
      <div style="height:16px"></div>
      <div class="row">
        <div class="node">
          <b>Redis</b>
          <span>nonce / ticket<br />限流计数</span>
        </div>
        <div class="arrow">⇄</div>
        <div class="node">
          <b>Kotlin sidecar</b>
          <span>事件流 → 规则</span>
        </div>
        <div class="arrow">⇄</div>
        <div class="node">
          <b>CrowdSec</b>
          <span>LAPI 决策</span>
        </div>
        <div class="arrow">⇄</div>
        <div class="node">
          <b>XDP/eBPF</b>
          <span>LPM 丢包</span>
        </div>
      </div>
    </div>
  </div>
</section>

<section id="config">
  <div class="wrap">
    <h2>单一配置源</h2>
    <p class="lead">所有服务器配置收敛到一个文件，渲染派生出容器、反向代理、策略与内核过滤。</p>
    <pre><code># config/config.yaml（节选）
server:
  host: 127.0.0.1
  port: 8080
  public_url: https://chat.example.com
  ws_allowed_origins: ["https://chat.example.com"]
jwt:
  secret: ${openssl rand -hex 32}
  ttl: 24h
crypto:
  pepper: ${openssl rand -hex 16}
  argon2: { memory_kib: 65536, iterations: 3, parallelism: 1 }
defense:
  rate_limits: { ws_per_ip: 32, prekey_per_min: 120, envelope_per_min: 240 }
  xdp: { enabled: true, iface: eth0 }
  crowdsec: { lapi_port: 8081 }
# make config  →  deploy/.env, caddy.env, acquis.yaml, application.yaml</code></pre>
  </div>
</section>

<section id="links">
  <div class="wrap">
    <h2>文档与路线</h2>
    <p class="lead">完整密码学承诺、自托管步骤与防御清单。</p>
    <div class="grid">
      <div class="card">
        <h3><a href="CRYPTO.md" style="color:var(--text)">密码学模型 →</a></h3>
        <p>账号派生链、X3DH 信封、Sealed Sender、零知识承诺表。</p>
      </div>
      <div class="card">
        <h3><a href="SELF_HOSTING.md" style="color:var(--text)">自托管指南 →</a></h3>
        <p>改 config、render、docker compose up、验证健康检查。</p>
      </div>
      <div class="card">
        <h3><a href="KNOWN_ERRORS.md" style="color:var(--text)">已知错误清单 →</a></h3>
        <p>A–J 协议层 + K-1–K-9 内核/策略层踩坑与防御。</p>
      </div>
    </div>
  </div>
</section>

<footer>
  <div class="wrap">
    <p>Paperplane · AGPL-3.0 · 图标：黑白折纸鸽子 · 部署在 GitHub Pages。</p>
  </div>
</footer>

</body>
</html>
```

### `docs/assets/logo.png`

_二进制资源（926208 B，sha256 `e2be9eac39cd665a690ad2597044f9f72d4e7014845697acf7625e023d2cb76f`）_。

### `policy-sidecar/Dockerfile`

```dockerfile
FROM gradle:8.10-jdk21 AS build
WORKDIR /app
COPY settings.gradle.kts build.gradle.kts ./
COPY src ./src
RUN gradle --no-daemon installDist

FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=build /app/build/install/paperplane-policy /app
ENV CONFIG_FILE=/app/config/application.yaml
ENTRYPOINT ["/app/bin/paperplane-policy"]
```

### `policy-sidecar/build.gradle.kts`

```kotlin
plugins {
    kotlin("jvm") version "2.0.21"
    application
}

group = "io.paperplane"
version = "0.1.0"

repositories { mavenCentral() }

dependencies {
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.9.0")
    implementation("io.lettuce:lettuce-core:6.5.2.RELEASE")
    implementation("org.yaml:snakeyaml:2.3")
    implementation("org.slf4j:slf4j-simple:2.0.16")
}

application {
    mainClass.set("io.paperplane.policy.MainKt")
}

tasks.withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile> {
    compilerOptions.freeCompilerArgs.addAll("-Xjsr305=strict")
}

tasks.jar {
    manifest.attributes["Main-Class"] = "io.paperplane.policy.MainKt"
}
```

### `policy-sidecar/settings.gradle.kts`

```kotlin
rootProject.name = "paperplane-policy"
```

### `policy-sidecar/config/application.yaml`

```yaml
# 由 render-config.sh 从 config/config.yaml 生成，请勿手改。
server:
  health_port: 8091
redis:
  host: "127.0.0.1"
  port: 6379
  password: "change-me-redis"
  database: 0
policy:
  stream: "paperplane:events"
  consumer_group: "paperplane-policy"
  consumer_name: "policy-1"
  window_seconds: 60
  ip_account_ratio: 8          # 同一 IP 关联账号数超阈值 -> 封禁 IP
  reconnect_limit: 5           # 窗口内重连次数超阈值 -> 降速/封禁
  reconnect_window_seconds: 30
  ban_seconds: 600
  fail_open_timeout_ms: 50
crowdsec:
  lapi_url: "http://127.0.0.1:8081"
  api_key_env: "CROWDSEC_LAPI_KEY"
  decision_duration: "10m"
```

### `policy-sidecar/src/main/kotlin/io/paperplane/policy/BanPublisher.kt`

```kotlin
package io.paperplane.policy

import io.lettuce.core.api.sync.RedisCommands
import io.lettuce.core.XAddArgs
import org.slf4j.LoggerFactory
import java.net.URI
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.http.HttpResponse
import java.time.Duration

/**
 * 下发封禁的唯一出口：
 *   1) Redis ZSET（Go 内嵌层 + XDP loader 消费）
 *   2) Redis Stream（/var/lib/paperplane/bans.txt 镜像）
 *   3) CrowdSec LAPI（机器人群封禁协同）
 *
 * 任何外部调用失败都 fail-open（50ms 超时），绝不阻塞事件消费。
 */
class BanPublisher(
    private val cfg: AppConfig,
    private val sync: RedisCommands<String, String>,
) {
    private val log = LoggerFactory.getLogger("bans")
    private val http = HttpClient.newBuilder()
        .connectTimeout(Duration.ofMillis(cfg.policy.failOpenTimeoutMs))
        .build()

    fun banIP(ip: String, reason: String) {
        val nowSec = redisNowSec()
        val until = nowSec + cfg.policy.banSeconds
        try {
            sync.zadd("paperplane:ban:ip", until.toDouble(), ip)
            sync.xadd(
                XAddArgs().stream("paperplane:bans").maxLen(10_000, true)
                    .add("scope", "ip").add("value", ip).add("reason", reason)
                    .add("until", until.toString())
            )
            crowdSec("ip", ip, reason)
        } catch (e: Exception) {
            log.warn("ban publish failed, failing open", e)
        }
    }

    fun banAccount(account: String, reason: String) {
        val until = redisNowSec() + cfg.policy.banSeconds
        try {
            sync.zadd("paperplane:ban:account", until.toDouble(), account)
            sync.xadd(
                XAddArgs().stream("paperplane:bans").maxLen(10_000, true)
                    .add("scope", "account").add("value", account).add("reason", reason)
                    .add("until", until.toString())
            )
        } catch (e: Exception) {
            log.warn("account ban publish failed", e)
        }
    }

    private fun crowdSec(scope: String, value: String, reason: String) {
        val apiKey = System.getenv(cfg.crowdsec.apiKeyEnv) ?: return
        if (apiKey.isBlank()) return
        try {
            val body = """{"duration":"${cfg.crowdsec.decisionDuration}","scenario":"paperplane:$reason","type":"$scope","value":"$value","origin":"policy-sidecar"}"""
            val req = HttpRequest.newBuilder()
                .uri(URI.create(cfg.crowdsec.lapiUrl.trimEnd('/') + "/v1/decisions"))
                .timeout(Duration.ofMillis(cfg.policy.failOpenTimeoutMs))
                .header("Authorization", "ApiKey $apiKey")
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(body))
                .build()
            http.send(req, HttpResponse.BodyHandlers.discarding())
        } catch (e: Exception) {
            log.debug("crowdsec call skipped", e)
        }
    }

    private fun redisNowSec(): Long = try {
        sync.time()[0].toLong()
    } catch (_: Exception) {
        System.currentTimeMillis() / 1000
    }
}
```

### `policy-sidecar/src/main/kotlin/io/paperplane/policy/Config.kt`

```kotlin
package io.paperplane.policy

import org.yaml.snakeyaml.Yaml
import java.io.File

/**
 * 单一配置源：所有字段都由 scripts/render-config.sh 从 config/config.yaml 生成。
 * 本类不允许出现硬编码阈值。
 */
data class AppConfig(
    val server: Server = Server(),
    val redis: Redis = Redis(),
    val policy: Policy = Policy(),
    val crowdsec: CrowdSec = CrowdSec(),
) {
    data class Server(val healthPort: Int = 8091)
    data class Redis(val host: String = "127.0.0.1", val port: Int = 6379,
                     val password: String = "", val db: Int = 0)
    data class Policy(
        val stream: String = "paperplane:events",
        val consumerGroup: String = "paperplane-policy",
        val consumerName: String = "policy-1",
        val windowSeconds: Long = 60,
        val ipAccountRatio: Int = 8,
        val reconnectLimit: Int = 5,
        val reconnectWindowSeconds: Long = 30,
        val banSeconds: Long = 600,
        val failOpenTimeoutMs: Long = 50,
    )
    data class CrowdSec(
        val lapiUrl: String = "http://127.0.0.1:8081",
        val apiKeyEnv: String = "CROWDSEC_LAPI_KEY",
        val decisionDuration: String = "10m",
    )

    companion object {
        @Suppress("UNCHECKED_CAST")
        fun load(path: String): AppConfig {
            val raw = Yaml().load<Map<String, Any>>(File(path).inputStream()) ?: emptyMap()
            fun section(name: String) = (raw[name] as? Map<String, Any>) ?: emptyMap()
            fun Map<String, Any>.str(k: String, d: String) = this[k] as? String ?: d
            fun Map<String, Any>.long(k: String, d: Long) = (this[k] as? Number)?.toLong() ?: d
            fun Map<String, Any>.int(k: String, d: Int) = (this[k] as? Number)?.toInt() ?: d

            val r = section("redis"); val p = section("policy"); val c = section("crowdsec")
            return AppConfig(
                redis = Redis(r.str("host", "127.0.0.1"), r.int("port", 6379),
                    r.str("password", ""), r.int("database", r.int("db", 0))),
                policy = Policy(
                    p.str("stream", "paperplane:events"),
                    p.str("consumer_group", "paperplane-policy"),
                    p.str("consumer_name", "policy-1"),
                    p.long("window_seconds", 60),
                    p.int("ip_account_ratio", 8),
                    p.int("reconnect_limit", 5),
                    p.long("reconnect_window_seconds", 30),
                    p.long("ban_seconds", 600),
                    p.long("fail_open_timeout_ms", 50),
                ),
                crowdsec = CrowdSec(
                    c.str("lapi_url", "http://127.0.0.1:8081"),
                    c.str("api_key_env", "CROWDSEC_LAPI_KEY"),
                    c.str("decision_duration", "10m"),
                ),
            )
        }
    }
}
```

### `policy-sidecar/src/main/kotlin/io/paperplane/policy/Main.kt`

```kotlin
package io.paperplane.policy

import io.lettuce.core.*
import kotlinx.coroutines.*
import org.slf4j.LoggerFactory
import java.time.Duration

/**
 * Kotlin sidecar 入口。
 * 消费 Redis Stream 事件流，运行跨维度封禁规则，结构化并发。
 */
suspend fun main(): Unit = coroutineScope {
    val log = LoggerFactory.getLogger("main")
    val cfgPath = System.getenv("CONFIG_FILE") ?: "config/application.yaml"
    val cfg = AppConfig.load(cfgPath)

    val redisUri = RedisURI.Builder
        .redis(cfg.redis.host, cfg.redis.port)
        .apply { if (cfg.redis.password.isNotEmpty()) withPassword(cfg.redis.password) }
        .withDatabase(cfg.redis.db)
        .withTimeout(Duration.ofMillis(cfg.policy.failOpenTimeoutMs + 100))
        .build()

    val client = RedisClient.create(redisUri)
    val conn = client.connect()
    val sync = conn.sync()

    // mkstream(true): 流不存在时自动创建；BUSYGROUP 异常吞掉。
    try {
        sync.xgroupCreate(
            XReadArgs.StreamOffset.latest(cfg.policy.stream),
            cfg.policy.consumerGroup,
            XGroupCreateArgs().mkstream(true),
        )
    } catch (e: RedisBusyException) {
        log.info("consumer group already exists")
    }

    val bans = BanPublisher(cfg, sync)
    val rules = Rules(cfg, sync, bans)

    log.info(
        "policy sidecar ready stream={} group={} consumer={}",
        cfg.policy.stream, cfg.policy.consumerGroup, cfg.policy.consumerName
    )

    while (isActive) {
        try {
            val msgs = sync.xreadgroup(
                Consumer.from(cfg.policy.consumerGroup, cfg.policy.consumerName),
                XReadArgs().count(50).block(Duration.ofSeconds(2)),
                XReadArgs.StreamOffset.lastConsumed(cfg.policy.stream),
            ) ?: emptyList()

            for (msg in msgs) {
                try {
                    // Lettuce 返回的 Map<key,value> 全部是 String。
                    rules.handle(msg.body)
                } catch (e: Exception) {
                    log.warn("rule failed on message {}", msg.id, e)
                }
                try {
                    conn.async().xack(cfg.policy.stream, cfg.policy.consumerGroup, msg.id)
                } catch (e: Exception) {
                    log.warn("xack failed {}", msg.id, e)
                }
            }
        } catch (e: Exception) {
            log.warn("stream read error, backoff", e)
            delay(2000)
        }
    }
}
```

### `policy-sidecar/src/main/kotlin/io/paperplane/policy/Rules.kt`

```kotlin
package io.paperplane.policy

import io.lettuce.core.api.sync.RedisCommands
import org.slf4j.LoggerFactory

/**
 * 跨维度规则引擎（纯协程结构化并发，每条消息独立异常隔离）。
 * 时钟统一来自 Redis TIME，避免本地时钟漂移（K-8）。
 */
class Rules(
    private val cfg: AppConfig,
    private val sync: RedisCommands<String, String>,
    private val bans: BanPublisher,
) {
    private val log = LoggerFactory.getLogger("rules")

    fun handle(fields: Map<String, String>) {
        val kind = fields["kind"] ?: return
        val ip = fields["ip"].orEmpty()
        val account = fields["account"].orEmpty()
        if (ip.isEmpty() && account.isEmpty()) return
        val nowMs = redisNowMs()
        when (kind) {
            "connect", "envelope" -> {
                if (ip.isNotEmpty() && account.isNotEmpty()) trackIpAccount(ip, account, nowMs)
                if (kind == "connect") trackReconnects(account)
            }
            "register", "auth_fail" -> {
                if (ip.isNotEmpty()) trackIpAccount(ip, account.ifEmpty { "anon" }, nowMs)
            }
        }
    }

    /** 同一 IP 短时聚合异常数量的账号：代理/批量注册特征。 */
    private fun trackIpAccount(ip: String, account: String, nowMs: Long) {
        val key = "paperplane:ipaccts:$ip"
        sync.zadd(key, nowMs.toDouble(), account)
        sync.zremrangebyscore(key, 0.0, (nowMs - cfg.policy.windowSeconds * 1000).toDouble())
        val uniq = sync.zcard(key)
        sync.expire(key, cfg.policy.windowSeconds + 10)
        if (uniq > cfg.policy.ipAccountRatio) {
            log.warn("ip_account_ratio exceeded ip={} uniq={}", ip, uniq)
            bans.banIP(ip, "ip_account_ratio")
        }
    }

    /** 同一账号高频重连：重连攻击/网络层抖动放大。 */
    private fun trackReconnects(account: String) {
        val key = "paperplane:acct_reconnect:$account"
        val n = sync.incr(key)
        sync.expire(key, cfg.policy.reconnectWindowSeconds)
        if (n > cfg.policy.reconnectLimit) {
            log.warn("reconnect storm account={} count={}", account, n)
            bans.banAccount(account, "reconnect_storm")
        }
    }

    /** Redis TIME -> epoch millis（K-8 统一时钟源）。 */
    private fun redisNowMs(): Long {
        return try {
            val t = sync.time()
            t[0].toLong() * 1000 + t[1].toLong() / 1000
        } catch (_: Exception) {
            System.currentTimeMillis()
        }
    }
}
```

### `scripts/gen-manifest.py`

```python
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
```

### `scripts/render-config.sh`

```bash
#!/usr/bin/env bash
# =============================================================================
# render-config.sh — 从唯一配置源 config/config.yaml 派生所有下游配置
# -----------------------------------------------------------------------------
# 用法:  ./scripts/render-config.sh [--check]
#   --check  只校验必要字段，不写任何文件（CI 用）
#
# 依赖:  yq (v4, 可选) 或 python3 + PyYAML
# 输出:  deploy/.env
#        deploy/caddy.env
#        deploy/crowdsec/acquis.yaml
#        policy-sidecar/config/application.yaml
#        xdp/config.env
# =============================================================================
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONFIG="${CONFIG_FILE:-$ROOT/config/config.yaml}"
CHECK_ONLY=0
[[ "${1:-}" == "--check" ]] && CHECK_ONLY=1

die()  { printf '\033[31m[render-config] %s\033[0m\n' "$*" >&2; exit 1; }
warn() { printf '\033[33m[render-config] %s\033[0m\n' "$*" >&2; }
info() { printf '\033[36m[render-config] %s\033[0m\n' "$*"; }

[[ -f "$CONFIG" ]] || die "找不到配置文件: $CONFIG"

if command -v yq >/dev/null 2>&1 && yq --version 2>&1 | grep -q 'yq.*4'; then
  RENDERER=yq
elif command -v python3 >/dev/null 2>&1 && python3 -c 'import yaml' >/dev/null 2>&1; then
  RENDERER=python
else
  die "需要 yq v4 或 python3 + PyYAML（pip install pyyaml）"
fi
info "配置源: $CONFIG (renderer=$RENDERER)"

# ---------------------------------------------------------------------------
# cfg <dotted.path>   读取单值（数组用逗号连接）
# cfg_list <path>     读取数组，每项一行
# ---------------------------------------------------------------------------
cfg() {
  local expr="$1"
  if [[ "$RENDERER" == yq ]]; then
    yq -r "$expr" "$CONFIG"
  else
    python3 - "$CONFIG" "$expr" <<'PY'
import sys, yaml
doc = yaml.safe_load(open(sys.argv[1], encoding="utf-8"))
cur = doc
for part in sys.argv[2].lstrip(".").split("."):
    if isinstance(cur, dict) and part in cur:
        cur = cur[part]
    else:
        cur = ""
        break
if isinstance(cur, bool):
    print("true" if cur else "false")
elif isinstance(cur, list):
    print(",".join(str(x) for x in cur))
elif cur is None:
    print("")
else:
    print(cur)
PY
  fi
}

cfg_list() {
  local path="$1"
  if [[ "$RENDERER" == yq ]]; then
    yq -r "$path[]" "$CONFIG"
  else
    python3 - "$CONFIG" "$path" <<'PY'
import sys, yaml
doc = yaml.safe_load(open(sys.argv[1], encoding="utf-8"))
cur = doc
for part in sys.argv[2].lstrip(".").split("."):
    cur = (cur or {}).get(part, [])
for item in (cur or []):
    print(item)
PY
  fi
}

require() {
  local name="$1" value="$2"
  [[ -n "$value" ]] || die "必填字段缺失: $name"
}

# yaml 双引号转义
yq_escape() { printf '%s' "${1//\\/\\\\}" | sed 's/"/\\"/g'; }

# "60s" / "1m" / "3600" -> 秒
to_seconds() {
  local v="$1"
  case "$v" in
    *s) printf '%s' "${v%s}" ;;
    *m) printf '%s' $(( ${v%m} * 60 )) ;;
    *h) printf '%s' $(( ${v%h} * 3600 )) ;;
    *)  printf '%s' "$v" ;;
  esac
}

PROJECT=$(cfg .project);                       require project "$PROJECT"
ENVIRONMENT=$(cfg .environment)
SERVER_HOST=$(cfg .server.host);               require server.host "$SERVER_HOST"
SERVER_PORT=$(cfg .server.port);               require server.port "$SERVER_PORT"
PUBLIC_URL=$(cfg .server.public_url);          require server.public_url "$PUBLIC_URL"
DB_HOST=$(cfg .database.host);                 require database.host "$DB_HOST"
DB_PORT=$(cfg .database.port)
DB_NAME=$(cfg .database.name);                 require database.name "$DB_NAME"
DB_USER=$(cfg .database.user);                 require database.user "$DB_USER"
DB_PASSWORD=$(cfg .database.password);         require database.password "$DB_PASSWORD"
REDIS_HOST=$(cfg .redis.host);                 require redis.host "$REDIS_HOST"
REDIS_PORT=$(cfg .redis.port)
REDIS_PASSWORD=$(cfg .redis.password);         require redis.password "$REDIS_PASSWORD"
JWT_SECRET=$(cfg .jwt.secret)
PEPPER=$(cfg .crypto.pepper)
LAPI_PORT=$(cfg .defense.crowdsec.lapi_port)
KOTLIN_STREAM=$(cfg .defense.kotlin.redis_stream)
KOTLIN_GROUP=$(cfg .defense.kotlin.consumer_group)
KOTLIN_WINDOW=$(to_seconds "$(cfg .defense.kotlin.window)")
WS_LIMIT=$(cfg .defense.rate_limits.ws_per_ip)
PREKEY_LIMIT=$(cfg .defense.rate_limits.prekey_per_min)
ENVELOPE_LIMIT=$(cfg .defense.rate_limits.envelope_per_min)
TARGET_LIMIT=$(cfg .defense.rate_limits.envelope_distinct_targets)
ENVELOPE_WINDOW=$(to_seconds "$(cfg .defense.rate_limits.envelope_window)")
CHALLENGE_LIMIT=$(cfg .defense.rate_limits.challenge_per_min)

[[ ${#JWT_SECRET} -ge 32 ]] || die "jwt.secret 至少 32 字节（当前 ${#JWT_SECRET}）"
[[ ${#PEPPER} -ge 16 ]]    || die "crypto.pepper 至少 16 字节（当前 ${#PEPPER}）"
[[ "$PUBLIC_URL" == https://* && "$ENVIRONMENT" == production ]] || warn "生产环境 public_url 必须是 https://"
[[ "$SERVER_HOST" == 127.0.0.1 || "$SERVER_HOST" == localhost ]] || warn "server.host 建议绑回环地址 127.0.0.1，公网流量交给 Caddy"

WS_ORIGINS=$(cfg_list .server.ws_allowed_origins | paste -sd, -)
[[ -n "$WS_ORIGINS" ]] || die "server.ws_allowed_origins 不能为空"

if grep -q 'change-me' <<<"$JWT_SECRET$PEPPER$DB_PASSWORD$REDIS_PASSWORD"; then
  if [[ "$ENVIRONMENT" == production ]]; then
    die "检测到占位密钥（change-me）。请在 config/config.yaml 里换成真实随机值后再部署。"
  fi
  warn "检测到占位密钥，仅 development 环境允许。"
fi

if [[ $CHECK_ONLY -eq 1 ]]; then
  info "校验通过：配置字段完整、长度合规。"
  exit 0
fi

mkdir -p "$ROOT/deploy" "$ROOT/deploy/crowdsec" "$ROOT/policy-sidecar/config" "$ROOT/xdp"

# ------------------------------- deploy/.env -------------------------------
cat > "$ROOT/deploy/.env" <<EOF
# 由 scripts/render-config.sh 从 config/config.yaml 生成，请勿手改。
COMPOSE_PROJECT_NAME=${PROJECT}
PUBLIC_URL=${PUBLIC_URL}
SERVER_HOST=${SERVER_HOST}
SERVER_PORT=${SERVER_PORT}
UPSTREAM=${SERVER_HOST}:${SERVER_PORT}
POSTGRES_DB=${DB_NAME}
POSTGRES_USER=${DB_USER}
POSTGRES_PASSWORD=${DB_PASSWORD}
POSTGRES_HOST=${DB_HOST}
POSTGRES_PORT=${DB_PORT}
REDIS_HOST=${REDIS_HOST}
REDIS_PORT=${REDIS_PORT}
REDIS_PASSWORD=${REDIS_PASSWORD}
WS_ALLOWED_ORIGINS=${WS_ORIGINS}
CONFIG_PATH=/etc/paperplane/config.yaml
CROWDSEC_LAPI_PORT=${LAPI_PORT}
KOTLIN_REDIS_STREAM=${KOTLIN_STREAM}
KOTLIN_CONSUMER_GROUP=${KOTLIN_GROUP}
KOTLIN_WINDOW_SECONDS=${KOTLIN_WINDOW}
WS_PER_IP=${WS_LIMIT}
PREKEY_PER_MIN=${PREKEY_LIMIT}
ENVELOPE_PER_MIN=${ENVELOPE_LIMIT}
ENVELOPE_DISTINCT_TARGETS=${TARGET_LIMIT}
ENVELOPE_WINDOW_SECONDS=${ENVELOPE_WINDOW}
CHALLENGE_PER_MIN=${CHALLENGE_LIMIT}
EOF

# ------------------------------ deploy/caddy.env ----------------------------
cat > "$ROOT/deploy/caddy.env" <<EOF
PUBLIC_URL=${PUBLIC_URL}
UPSTREAM=${SERVER_HOST}:${SERVER_PORT}
EOF

# -------------------------- crowdsec/acquis.yaml ---------------------------
cat > "$ROOT/deploy/crowdsec/acquis.yaml" <<EOF
# 由 render-config.sh 生成：Caddy 访问日志 -> CrowdSec 场景引擎
filenames:
  - /var/log/caddy/access.log
labels:
  type: caddy
EOF

# ------------------- policy-sidecar/config/application.yaml ----------------
cat > "$ROOT/policy-sidecar/config/application.yaml" <<EOF
# 由 render-config.sh 从 config/config.yaml 生成，请勿手改。
server:
  health_port: 8091
redis:
  host: "$(yq_escape "$REDIS_HOST")"
  port: ${REDIS_PORT}
  password: "$(yq_escape "$REDIS_PASSWORD")"
  database: $(cfg .redis.db)
policy:
  stream: "$(yq_escape "$KOTLIN_STREAM")"
  consumer_group: "$(yq_escape "$KOTLIN_GROUP")"
  consumer_name: "policy-1"
  window_seconds: ${KOTLIN_WINDOW}
  ip_account_ratio: 8          # 同一 IP 关联账号数超阈值 -> 封禁 IP
  reconnect_limit: 5           # 窗口内重连次数超阈值 -> 降速/封禁
  reconnect_window_seconds: 30
  ban_seconds: 600
  fail_open_timeout_ms: 50
crowdsec:
  lapi_url: "http://127.0.0.1:${LAPI_PORT}"
  api_key_env: "CROWDSEC_LAPI_KEY"
  decision_duration: "10m"
EOF

# ------------------------------ xdp/config.env ------------------------------
cat > "$ROOT/xdp/config.env" <<EOF
XDP_ENABLED=$(cfg .defense.xdp.enabled)
XDP_IFACE=$(cfg .defense.xdp.iface)
BAN_FILE=/var/lib/paperplane/bans.txt
BAN_POLL_SECONDS=5
EOF

info "已生成："
info "  deploy/.env"
info "  deploy/caddy.env"
info "  deploy/crowdsec/acquis.yaml"
info "  policy-sidecar/config/application.yaml"
info "  xdp/config.env"
```

### `scripts/termux-bootstrap.sh`

```bash
#!/data/data/com.termux/files/usr/bin/env bash
# Paperplane 在 Termux 里直接运行（服务端）。
# 客户端仍需在 Android 上以 APK 形式安装；这里跑的是 Go 服务端 + PG/Redis。
set -euo pipefail

echo "==> 安装依赖（golang / postgresql / redis / git）"
pkg update -y
pkg install -y golang git postgresql redis clang make

export PREFIX=${PREFIX:-/data/data/com.termux/files/usr}

echo "==> 初始化 PostgreSQL"
if [ ! -d "$PREFIX/var/lib/postgresql" ]; then
  initdb "$PREFIX/var/lib/postgresql"
fi
pg_ctl -D "$PREFIX/var/lib/postgresql" -l "$PREFIX/var/log/postgresql.log" start || true
sleep 1
createuser -s paperplane 2>/dev/null || true
createdb paperplane 2>/dev/null || true

echo "==> 启动 Redis"
redis-server --daemonize yes --port 6379

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "==> 启动 Paperplane 服务端"
exec go run ./server/cmd/server -config config/config.yaml
```

### `server/Dockerfile`

```dockerfile
# Paperplane Go 服务端镜像
# 构建上下文必须为 server/ 目录
FROM golang:1.25-alpine AS build
WORKDIR /src
RUN apk add --no-cache git ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags "-s -w" -o /out/paperplane-server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/paperplane-server /app/paperplane-server
COPY migrations /migrations
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app/paperplane-server"]
```

### `server/go.mod`

```text
module github.com/paperplane/paperplane/server

go 1.25.0

require (
	github.com/gorilla/websocket v1.5.3
	github.com/jackc/pgx/v5 v5.7.6
	github.com/redis/go-redis/v9 v9.7.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)
```

### `server/go.sum`

```text
github.com/bsm/ginkgo/v2 v2.12.0 h1:Ny8MWAHyOepLGlLKYmXG4IEkioBysk6GpaRTLC8zwWs=
github.com/bsm/ginkgo/v2 v2.12.0/go.mod h1:SwYbGRRDovPVboqFv0tPTcG1sN61LM1Z4ARdbAV9g4c=
github.com/bsm/gomega v1.27.10 h1:yeMWxP2pV2fG3FgAODIY8EiRE3dy0aeFYt4l7wh6yKA=
github.com/bsm/gomega v1.27.10/go.mod h1:JyEr/xRbxbtgWNi8tIEVPUYZ5Dzef52k01W3YH0H+O0=
github.com/cespare/xxhash/v2 v2.2.0 h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=
github.com/cespare/xxhash/v2 v2.2.0/go.mod h1:VGX0DQ3Q6kWi7AoAeZDth3/j3BFtOZR5XLFGgcrjCOs=
github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=
github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f/go.mod h1:cuUVRXasLTGF7a8hSLbxyZXjz+1KgoB3wDUb6vlszIc=
github.com/gorilla/websocket v1.5.3 h1:saDtZ6Pbx/0u+bgYQ3q96pZgCzfhKXGPqt7kZ72aNNg=
github.com/gorilla/websocket v1.5.3/go.mod h1:YR8l580nyteQvAITg2hZ9XVh4b55+EU/adAjf1fMHhE=
github.com/jackc/pgpassfile v1.0.0 h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=
github.com/jackc/pgpassfile v1.0.0/go.mod h1:CEx0iS5ambNFdcRtxPj5JhEz+xB6uRky5eyVu/W2HEg=
github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 h1:iCEnooe7UlwOQYpKFhBabPMi4aNAfoODPEFNiAnClxo=
github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761/go.mod h1:5TJZWKEWniPve33vlWYSoGYefn3gLQRzjfDlhSJ9ZKM=
github.com/jackc/pgx/v5 v5.7.6 h1:rWQc5FwZSPX58r1OQmkuaNicxdmExaEz5A2DO2hUuTk=
github.com/jackc/pgx/v5 v5.7.6/go.mod h1:aruU7o91Tc2q2cFp5h4uP3f6ztExVpyVv88Xl/8Vl8M=
github.com/jackc/puddle/v2 v2.2.2 h1:PR8nw+E/1w0GLuRFSmiioY6UooMp6KJv0/61nB7icHo=
github.com/jackc/puddle/v2 v2.2.2/go.mod h1:vriiEXHvEE654aYKXXjOvZM39qJ0q+azkZFrfEOc3H4=
github.com/kr/pretty v0.3.0 h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=
github.com/kr/pretty v0.3.0/go.mod h1:640gp4NfQd8pI5XOwp5fnNeVWj67G7CFk/SaSQn7NBk=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/redis/go-redis/v9 v9.7.0 h1:HhLSs+B6O021gwzl+locl0zEDnyNkxMtf/Z3NNBMa9E=
github.com/redis/go-redis/v9 v9.7.0/go.mod h1:f6zhXITC7JUJIlPEiBOTXxJgPLdZcA93GewI7inzyWw=
github.com/rogpeppe/go-internal v1.16.0 h1:O9DK+vNMDVGLr2BeZqmpLeMjiMNkuXfcqntWbZV6S5g=
github.com/rogpeppe/go-internal v1.16.0/go.mod h1:DrUVZyrJU+txYW5/1kwtXQSMFio52ZOxX7yM1VHvnxs=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.1 h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=
github.com/stretchr/testify v1.8.1/go.mod h1:w2LPCIKwWwSfY2zedu0+kehJoqGctiVI29o6fzry7u4=
golang.org/x/crypto v0.37.0 h1:kJNSjF/Xp7kU0iB2Z+9viTPMW4EqqsrywMXLJOOsXSE=
golang.org/x/crypto v0.37.0/go.mod h1:vg+k43peMZ0pUMhYmVAWysMK35e6ioLh3wB8ZCAfbVc=
golang.org/x/sync v0.13.0 h1:AauUjRAJ9OSnvULf/ARrrVywoJDy0YS2AwQ98I37610=
golang.org/x/sync v0.13.0/go.mod h1:1dzgHSNfp02xaA81J2MS99Qcpr2w7fw1gpm99rleRqA=
golang.org/x/text v0.24.0 h1:dd5Bzh4yt5KYA8f9CJHCP4FB4D51c2c6JvN37xJJkJ0=
golang.org/x/text v0.24.0/go.mod h1:L8rBsPeo2pSS+xqN0d5u2ikmjtmoJbDBT1b7nHvFCdU=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

### `server/cmd/server/main.go`

```go
// Command server is the Paperplane relay. It terminates no TLS itself (Caddy
// does), decrypts nothing, and keeps ciphertext only until delivery is
// acknowledged.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/paperplane/paperplane/server/internal/auth"
	"github.com/paperplane/paperplane/server/internal/config"
	"github.com/paperplane/paperplane/server/internal/defense"
	"github.com/paperplane/paperplane/server/internal/httpapi"
	"github.com/paperplane/paperplane/server/internal/store"
	"github.com/paperplane/paperplane/server/internal/ws"
)

func main() {
	if err := run(); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run() error {
	configPath := flag.String("config", envOr("CONFIG_PATH", "config/config.yaml"), "path to config.yaml")
	migrationsDir := flag.String("migrations", envOr("MIGRATIONS_DIR", "migrations"), "directory with .sql migrations")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg, err := config.Load(*configPath)
	if err != nil {
		return err // fail-fast: never start with a half-valid configuration
	}
	if cfg.IsProduction() && secretLooksLikePlaceholder(cfg) {
		return errors.New("refusing to start production with placeholder secrets in config.yaml")
	}
	logger.Info("config loaded",
		"project", cfg.Project,
		"environment", cfg.Environment,
		"listen", cfg.ListenAddr(),
		"public_url", cfg.Server.PublicURL,
	)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := waitPostgres(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defer db.Close()

	migrations, err := loadMigrations(*migrationsDir)
	if err != nil {
		return err
	}
	for name, script := range migrations {
		if err := db.Migrate(ctx, script); err != nil {
			return fmt.Errorf("migration %s: %w", name, err)
		}
		logger.Info("migration applied", "file", name)
	}

	rdb := store.NewRedis(cfg)
	defer func() { _ = rdb.Close() }()
	if err := waitRedis(ctx, func(c context.Context) error { return rdb.Ping(c).Err() }, logger); err != nil {
		return err
	}

	limiter := defense.NewLimiter(rdb, cfg.Defense.Kotlin.RedisStream, logger)
	hub := ws.NewHub(logger)
	jwtIssuer := auth.NewIssuer(cfg.JWT.Secret, cfg.JWTTTL(), cfg.JWT.Issuer)
	api := httpapi.New(cfg, db, rdb, limiter, hub, jwtIssuer, logger)

	srv := &http.Server{
		Addr:              cfg.ListenAddr(),
		Handler:           api.Handler(),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout(),
		IdleTimeout:       cfg.IdleTimeout(),
		MaxHeaderBytes:    1 << 20,
		ErrorLog:          slog.NewLogLogger(logger.Handler(), slog.LevelWarn),
	}

	serveErr := make(chan error, 1)
	go func() {
		logger.Info("http listening", "addr", cfg.ListenAddr())
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
	}()

	// Kernel-space ban mirror for the XDP loader.
	go defense.SyncBanFile(ctx, rdb, envOr("BAN_FILE", "/var/lib/paperplane/bans.txt"),
		5*time.Second, logger)
	go cleanupLoop(ctx, db, logger)

	select {
	case err := <-serveErr:
		return fmt.Errorf("http server: %w", err)
	case <-ctx.Done():
	}

	logger.Info("shutdown requested")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout())
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Warn("graceful shutdown incomplete", "err", err)
	}
	hub.CloseAll()
	logger.Info("bye")
	return nil
}

func waitPostgres(ctx context.Context, cfg *config.Config, log *slog.Logger) (*store.Postgres, error) {
	var lastErr error
	for attempt := 1; attempt <= 20; attempt++ {
		db, err := store.NewPostgres(ctx, cfg.PostgresDSN())
		if err == nil {
			pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
			err = db.Ping(pingCtx)
			cancel()
			if err == nil {
				log.Info("postgres connected", "attempt", attempt)
				return db, nil
			}
			db.Close()
		}
		lastErr = err
		log.Warn("postgres not ready, retrying", "attempt", attempt, "err", err)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return nil, fmt.Errorf("postgres unreachable: %w", lastErr)
}

func waitRedis(ctx context.Context, ping func(context.Context) error, log *slog.Logger) error {
	var lastErr error
	for attempt := 1; attempt <= 20; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		err := ping(pingCtx)
		cancel()
		if err == nil {
			log.Info("redis connected", "attempt", attempt)
			return nil
		}
		lastErr = err
		log.Warn("redis not ready, retrying", "attempt", attempt, "err", err)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return fmt.Errorf("redis unreachable: %w", lastErr)
}

func cleanupLoop(ctx context.Context, db *store.Postgres, log *slog.Logger) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			runCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			n, err := db.CleanupExpiredEnvelopes(runCtx)
			cancel()
			if err != nil {
				log.Warn("envelope cleanup failed", "err", err)
				continue
			}
			if n > 0 {
				log.Info("expired envelopes removed", "count", n)
			}
		}
	}
}

func loadMigrations(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read migrations dir %q: %w", dir, err)
	}
	out := map[string]string{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		// Map iteration is random; migrations here are idempotent and ordered
		// by a single file, so name keying is enough for this project.
		out[e.Name()] = string(raw)
	}
	if len(out) == 0 {
		return nil, errors.New("no .sql migrations found")
	}
	return out, nil
}

func secretLooksLikePlaceholder(cfg *config.Config) bool {
	joined := cfg.JWT.Secret + cfg.Crypto.Pepper + cfg.Database.Password + cfg.Redis.Password
	return strings.Contains(joined, "change-me")
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
```

### `server/internal/auth/token.go`

```go
// Package auth implements the zero-knowledge account model and HS256 tokens.
//
// The server never sees a password. The client derives:
//
//	account_id = SHA-256(username || pepper)
//	argon2id   = Argon2id(password, salt = account_id[:16], m=64MiB, t=3, p=1)
//
// and sends only account_id (implicitly, by username) plus the PHC string.
// At login the client proves possession of that hash by HMAC-ing a single-use
// server nonce, so a captured credential cannot be replayed.
package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

const phcPrefix = "$argon2id$v=19$"

// base64url without padding keeps tokens URL/query safe.
var b64 = base64.RawURLEncoding

// UsernamePattern is deliberately small and ASCII-only: account ids are
// derived from it, so ambiguity would create duplicate-looking identities.
var UsernamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9._-]{2,31}$`)

// NormalizeUsername lowercases and trims; the result is what gets hashed.
func NormalizeUsername(raw string) string {
	return strings.ToLower(strings.TrimSpace(raw))
}

func ValidUsername(normalized string) bool { return UsernamePattern.MatchString(normalized) }

// AccountID is the irreversible account handle stored by the server.
func AccountID(pepper, normalizedUsername string) string {
	sum := sha256.Sum256([]byte(normalizedUsername + pepper))
	return hex.EncodeToString(sum[:])
}

// NewNonce returns 32 random bytes, base64url encoded.
func NewNonce() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return b64.EncodeToString(b), nil
}

// NewTicket returns a single-use WebSocket ticket.
func NewTicket() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return b64.EncodeToString(b), nil
}

// Proof binds the stored credential to a server nonce.
func Proof(credential, nonce string) string {
	mac := hmac.New(sha256.New, []byte(credential))
	mac.Write([]byte(nonce))
	return b64.EncodeToString(mac.Sum(nil))
}

func VerifyProof(credential, nonce, proof string) bool {
	expected := Proof(credential, nonce)
	return subtle.ConstantTimeCompare([]byte(expected), []byte(proof)) == 1
}

// ParsePHC validates the Argon2id PHC string shape and extracts m, t, p.
// The server rejects credentials weaker than the configured parameters.
func ParsePHC(encoded string) (memoryKiB, iterations uint32, parallelism uint8, err error) {
	if !strings.HasPrefix(encoded, phcPrefix) {
		return 0, 0, 0, errors.New("credential must be an argon2id v=19 PHC string")
	}
	rest := strings.TrimPrefix(encoded, phcPrefix)
	parts := strings.Split(rest, "$")
	if len(parts) != 3 {
		return 0, 0, 0, errors.New("malformed PHC string")
	}
	for _, kv := range strings.Split(parts[0], ",") {
		k, v, ok := strings.Cut(kv, "=")
		if !ok {
			return 0, 0, 0, errors.New("malformed argon2 parameters")
		}
		n, convErr := strconv.ParseUint(v, 10, 32)
		if convErr != nil {
			return 0, 0, 0, fmt.Errorf("bad argon2 parameter %q", kv)
		}
		switch k {
		case "m":
			memoryKiB = uint32(n)
		case "t":
			iterations = uint32(n)
		case "p":
			if n > 255 {
				return 0, 0, 0, errors.New("argon2 parallelism out of range")
			}
			parallelism = uint8(n)
		}
	}
	if memoryKiB == 0 || iterations == 0 || parallelism == 0 {
		return 0, 0, 0, errors.New("incomplete argon2 parameters")
	}
	if len(parts[1]) < 8 || len(parts[2]) < 16 {
		return 0, 0, 0, errors.New("argon2 salt/hash too short")
	}
	return memoryKiB, iterations, parallelism, nil
}

// Issuer creates and verifies HS256 JSON Web Tokens containing only account_id.
type Issuer struct {
	secret []byte
	ttl    time.Duration
	issuer string
	now    func() time.Time
}

type claims struct {
	Sub string `json:"sub"`
	Iss string `json:"iss"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

func NewIssuer(secret string, ttl time.Duration, issuer string) *Issuer {
	return &Issuer{secret: []byte(secret), ttl: ttl, issuer: issuer, now: time.Now}
}

// Issue returns a signed token. Header is fixed: {"alg":"HS256","typ":"JWT"}.
func (i *Issuer) Issue(accountID string) (string, error) {
	now := i.now()
	c := claims{Sub: accountID, Iss: i.issuer, Iat: now.Unix(), Exp: now.Add(i.ttl).Unix()}
	payload, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	header := b64.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	body := header + "." + b64.EncodeToString(payload)
	sig := i.sign(body)
	return body + "." + b64.EncodeToString(sig), nil
}

func (i *Issuer) sign(body string) []byte {
	mac := hmac.New(sha256.New, i.secret)
	mac.Write([]byte(body))
	return mac.Sum(nil)
}

// Verify checks structure, algorithm, signature, issuer and expiry.
func (i *Issuer) Verify(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", ErrInvalidToken
	}
	headerJSON, err := b64.DecodeString(parts[0])
	if err != nil {
		return "", ErrInvalidToken
	}
	var h struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	if err := json.Unmarshal(headerJSON, &h); err != nil || h.Alg != "HS256" {
		return "", ErrInvalidToken
	}
	expected := i.sign(parts[0] + "." + parts[1])
	got, err := b64.DecodeString(parts[2])
	if err != nil || subtle.ConstantTimeCompare(expected, got) != 1 {
		return "", ErrInvalidToken
	}
	payload, err := b64.DecodeString(parts[1])
	if err != nil {
		return "", ErrInvalidToken
	}
	var c claims
	if err := json.Unmarshal(payload, &c); err != nil {
		return "", ErrInvalidToken
	}
	if c.Iss != i.issuer || c.Sub == "" {
		return "", ErrInvalidToken
	}
	if i.now().Unix() >= c.Exp {
		return "", ErrExpiredToken
	}
	return c.Sub, nil
}

// TTL exposes the configured lifetime for API responses.
func (i *Issuer) TTL() time.Duration { return i.ttl }
```

### `server/internal/config/config.go`

```go
// Package config loads and validates the single source of truth:
// config/config.yaml. Every service reads the same file; nothing else is
// allowed to carry its own copy of server settings.
package config

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Config mirrors config/config.yaml field for field.
type Config struct {
	Project     string   `yaml:"project"`
	Environment string   `yaml:"environment"`
	Server      Server   `yaml:"server"`
	Database    Database `yaml:"database"`
	Redis       Redis    `yaml:"redis"`
	JWT         JWT      `yaml:"jwt"`
	Crypto      Crypto   `yaml:"crypto"`
	Client      Client   `yaml:"client"`
	Defense     Defense  `yaml:"defense"`
}

type Server struct {
	Host              string   `yaml:"host"`
	Port              int      `yaml:"port"`
	PublicURL         string   `yaml:"public_url"`
	WSAllowedOrigins  []string `yaml:"ws_allowed_origins"`
	ReadHeaderTimeout string   `yaml:"read_header_timeout"`
	IdleTimeout       string   `yaml:"idle_timeout"`
	ShutdownTimeout   string   `yaml:"shutdown_timeout"`
}

type Database struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Name            string `yaml:"name"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	MaxConns        int32  `yaml:"max_conns"`
	EnvelopeTTLHour int    `yaml:"envelope_ttl_hours"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	TTL    string `yaml:"ttl"`
	Issuer string `yaml:"issuer"`
}

type Crypto struct {
	Pepper string `yaml:"pepper"`
	Argon2 struct {
		MemoryKiB   uint32 `yaml:"memory_kib"`
		Iterations  uint32 `yaml:"iterations"`
		Parallelism uint8  `yaml:"parallelism"`
	} `yaml:"argon2"`
}

type Client struct {
	DevServerURL string `yaml:"dev_server_url"`
	AppName      string `yaml:"app_name"`
}

type Defense struct {
	XDP        XDP        `yaml:"xdp"`
	RateLimits RateLimits `yaml:"rate_limits"`
	Kotlin     Kotlin     `yaml:"kotlin"`
	CrowdSec   CrowdSec   `yaml:"crowdsec"`
}

type XDP struct {
	Enabled bool   `yaml:"enabled"`
	Iface   string `yaml:"iface"`
}

type RateLimits struct {
	WSPerIP                 int    `yaml:"ws_per_ip"`
	PreKeyPerMin            int    `yaml:"prekey_per_min"`
	EnvelopePerMin          int    `yaml:"envelope_per_min"`
	EnvelopeDistinctTargets int    `yaml:"envelope_distinct_targets"`
	EnvelopeWindow          string `yaml:"envelope_window"`
	ChallengePerMin         int    `yaml:"challenge_per_min"`
}

type Kotlin struct {
	RedisStream   string `yaml:"redis_stream"`
	ConsumerGroup string `yaml:"consumer_group"`
	Window        string `yaml:"window"`
}

type CrowdSec struct {
	Collections []string `yaml:"collections"`
	LAPIPort    int      `yaml:"lapi_port"`
}

// Load reads config.yaml, rejects unknown keys and fails fast on invalid input.
func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}
	var c Config
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true) // typo in config.yaml is a startup error, not a silent default
	if err := dec.Decode(&c); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}
	c.applyDefaults()
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Config) applyDefaults() {
	if c.Server.ReadHeaderTimeout == "" {
		c.Server.ReadHeaderTimeout = "10s"
	}
	if c.Server.IdleTimeout == "" {
		c.Server.IdleTimeout = "120s"
	}
	if c.Server.ShutdownTimeout == "" {
		c.Server.ShutdownTimeout = "15s"
	}
	if c.JWT.TTL == "" {
		c.JWT.TTL = "24h"
	}
	if c.JWT.Issuer == "" {
		c.JWT.Issuer = c.Project
	}
	if c.Database.MaxConns == 0 {
		c.Database.MaxConns = 16
	}
	if c.Database.EnvelopeTTLHour == 0 {
		c.Database.EnvelopeTTLHour = 24
	}
	if c.Defense.RateLimits.EnvelopeWindow == "" {
		c.Defense.RateLimits.EnvelopeWindow = "60s"
	}
	if c.Defense.Kotlin.RedisStream == "" {
		c.Defense.Kotlin.RedisStream = "paperplane:events"
	}
	if c.Defense.Kotlin.ConsumerGroup == "" {
		c.Defense.Kotlin.ConsumerGroup = "paperplane-policy"
	}
	if c.Defense.Kotlin.Window == "" {
		c.Defense.Kotlin.Window = "60s"
	}
	if c.Defense.CrowdSec.LAPIPort == 0 {
		c.Defense.CrowdSec.LAPIPort = 8081
	}
}

// Validate enforces every field the runtime depends on. Missing or nonsensical
// values abort startup instead of surfacing as a confusing runtime bug.
func (c *Config) Validate() error {
	problems := make([]string, 0, 8)

	if strings.TrimSpace(c.Project) == "" {
		problems = append(problems, "project is required")
	}
	switch c.Environment {
	case "production", "development":
	default:
		problems = append(problems, `environment must be "production" or "development"`)
	}
	if c.Server.Host == "" {
		problems = append(problems, "server.host is required")
	}
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		problems = append(problems, "server.port must be 1-65535")
	}
	u, err := url.Parse(c.Server.PublicURL)
	if err != nil || u.Host == "" {
		problems = append(problems, "server.public_url must be an absolute URL")
	} else if c.Environment == "production" && u.Scheme != "https" {
		problems = append(problems, "server.public_url must use https:// in production")
	}
	if len(c.Server.WSAllowedOrigins) == 0 {
		problems = append(problems, "server.ws_allowed_origins must not be empty")
	}
	for _, o := range c.Server.WSAllowedOrigins {
		if _, err := url.Parse(o); err != nil || !strings.Contains(o, "://") {
			problems = append(problems, "invalid ws origin: "+o)
		}
	}
	if c.Database.Host == "" || c.Database.Name == "" || c.Database.User == "" {
		problems = append(problems, "database host/name/user are required")
	}
	if c.Database.Port < 1 || c.Database.Port > 65535 {
		problems = append(problems, "database.port must be 1-65535")
	}
	if c.Database.Password == "" {
		problems = append(problems, "database.password is required")
	}
	if c.Redis.Host == "" {
		problems = append(problems, "redis.host is required")
	}
	if c.Redis.Port < 1 || c.Redis.Port > 65535 {
		problems = append(problems, "redis.port must be 1-65535")
	}
	if len(c.JWT.Secret) < 32 {
		problems = append(problems, "jwt.secret must be at least 32 bytes")
	}
	if len(c.Crypto.Pepper) < 16 {
		problems = append(problems, "crypto.pepper must be at least 16 bytes")
	}
	if c.Crypto.Argon2.MemoryKiB < 19456 {
		problems = append(problems, "crypto.argon2.memory_kib must be >= 19456 (19 MiB)")
	}
	if c.Crypto.Argon2.Iterations < 1 || c.Crypto.Argon2.Iterations > 10 {
		problems = append(problems, "crypto.argon2.iterations must be 1-10")
	}
	if c.Crypto.Argon2.Parallelism < 1 {
		problems = append(problems, "crypto.argon2.parallelism must be >= 1")
	}
	if c.Defense.RateLimits.WSPerIP < 1 || c.Defense.RateLimits.PreKeyPerMin < 1 ||
		c.Defense.RateLimits.EnvelopePerMin < 1 || c.Defense.RateLimits.ChallengePerMin < 1 {
		problems = append(problems, "defense.rate_limits values must be positive")
	}
	if c.Defense.XDP.Enabled && c.Defense.XDP.Iface == "" {
		problems = append(problems, "defense.xdp.iface is required when xdp is enabled")
	}
	if _, err := time.ParseDuration(c.JWT.TTL); err != nil {
		problems = append(problems, "jwt.ttl is not a duration (e.g. 24h)")
	}
	for _, d := range []struct{ name, value string }{
		{"server.read_header_timeout", c.Server.ReadHeaderTimeout},
		{"server.idle_timeout", c.Server.IdleTimeout},
		{"server.shutdown_timeout", c.Server.ShutdownTimeout},
		{"defense.rate_limits.envelope_window", c.Defense.RateLimits.EnvelopeWindow},
		{"defense.kotlin.window", c.Defense.Kotlin.Window},
	} {
		if _, err := time.ParseDuration(d.value); err != nil {
			problems = append(problems, d.name+" is not a duration")
		}
	}
	if len(problems) > 0 {
		return fmt.Errorf("invalid config: %s", strings.Join(problems, "; "))
	}
	return nil
}

func (c *Config) IsProduction() bool { return c.Environment == "production" }

func (c *Config) ListenAddr() string {
	return net.JoinHostPort(c.Server.Host, strconv.Itoa(c.Server.Port))
}

func (c *Config) PostgresDSN() string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.Database.User, c.Database.Password),
		Host:   net.JoinHostPort(c.Database.Host, strconv.Itoa(c.Database.Port)),
		Path:   "/" + c.Database.Name,
	}
	q := u.Query()
	q.Set("sslmode", "disable")
	q.Set("pool_max_conns", strconv.Itoa(int(c.Database.MaxConns)))
	q.Set("application_name", c.Project+"-server")
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Config) RedisAddr() string {
	return net.JoinHostPort(c.Redis.Host, strconv.Itoa(c.Redis.Port))
}

func (c *Config) JWTTTL() time.Duration { return mustDuration(c.JWT.TTL, 24*time.Hour) }
func (c *Config) ReadHeaderTimeout() time.Duration {
	return mustDuration(c.Server.ReadHeaderTimeout, 10*time.Second)
}
func (c *Config) IdleTimeout() time.Duration {
	return mustDuration(c.Server.IdleTimeout, 2*time.Minute)
}
func (c *Config) ShutdownTimeout() time.Duration {
	return mustDuration(c.Server.ShutdownTimeout, 15*time.Second)
}
func (c *Config) EnvelopeWindow() time.Duration {
	return mustDuration(c.Defense.RateLimits.EnvelopeWindow, time.Minute)
}
func (c *Config) PolicyWindow() time.Duration {
	return mustDuration(c.Defense.Kotlin.Window, time.Minute)
}

func mustDuration(s string, fallback time.Duration) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return fallback
	}
	return d
}
```

### `server/internal/defense/bans.go`

```go
package defense

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// SyncBanFile mirrors the Redis ban ZSET to a plain text file that the XDP
// loader watches. Kernel-space filtering therefore keeps working even when the
// userspace loader and Redis are briefly out of sync; the file is the only
// contract between Go and C.
//
// Format: one "ip expiry_unix" pair per line.
func SyncBanFile(ctx context.Context, rdb *redis.Client, path string, every time.Duration, log *slog.Logger) {
	if path == "" {
		return
	}
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	var lastErr string
	write := func() {
		now := time.Now().Unix()
		ips, err := rdb.ZRangeByScore(ctx, banKey, &redis.ZRangeBy{
			Min: "(" + strconv.FormatInt(now, 10),
			Max: "+inf",
		}).Result()
		if err != nil {
			log.Warn("ban sync read failed", "err", err)
			return
		}
		var b strings.Builder
		for _, ip := range ips {
			score, err := rdb.ZScore(ctx, banKey, ip).Result()
			if err != nil {
				continue
			}
			fmt.Fprintf(&b, "%s %d\n", ip, int64(score))
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			if lastErr != err.Error() {
				log.Warn("ban file directory unavailable", "path", path, "err", err)
				lastErr = err.Error()
			}
			return
		}
		tmp := path + ".tmp"
		if err := os.WriteFile(tmp, []byte(b.String()), 0o644); err != nil {
			if lastErr != err.Error() {
				log.Warn("ban file write failed", "err", err)
				lastErr = err.Error()
			}
			return
		}
		if err := os.Rename(tmp, path); err != nil {
			log.Warn("ban file rename failed", "err", err)
			return
		}
		lastErr = ""
	}

	write()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			write()
		}
	}
}
```

### `server/internal/defense/limits.go`

```go
// Package defense implements the embedded, in-process defence layer:
// token buckets, distinct-target sliding windows, ban lookups and the bridge
// that feeds kernel-space bans to the XDP loader.
//
// Every Redis call in the hot path is bounded and fails open. A Redis outage
// must degrade protection, never take the chat service down.
package defense

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const banKey = "paperplane:ban:ip"

type Limiter struct {
	rdb    *redis.Client
	log    *slog.Logger
	stream string
}

func NewLimiter(rdb *redis.Client, stream string, log *slog.Logger) *Limiter {
	return &Limiter{rdb: rdb, log: log, stream: stream}
}

var tokenBucketScript = redis.NewScript(`
local key = KEYS[1]
local rate = tonumber(ARGV[1])
local burst = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local pair = redis.call('HMGET', key, 'tokens', 'ts')
local tokens = tonumber(pair[1])
local ts = tonumber(pair[2])
if tokens == nil then tokens = burst; ts = now end
local elapsed = now - ts
if elapsed < 0 then elapsed = 0 end
tokens = math.min(burst, tokens + elapsed * rate / 60)
local allowed = 0
if tokens >= 1 then
  tokens = tokens - 1
  allowed = 1
end
redis.call('HSET', key, 'tokens', tokens, 'ts', now)
redis.call('EXPIRE', key, math.ceil(burst * 60 / rate) + 60)
return allowed
`)

// AllowTokenBucket consumes one token for subject in the given bucket.
// Fail-open on Redis errors (logged), so an outage cannot lock users out.
func (l *Limiter) AllowTokenBucket(ctx context.Context, bucket, subject string, perMinute int) bool {
	if perMinute <= 0 {
		return true
	}
	key := "paperplane:rl:" + bucket + ":" + subject
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	n, err := tokenBucketScript.Run(ctx, l.rdb, []string{key},
		perMinute, perMinute, time.Now().Unix()).Int()
	if err != nil {
		l.log.Warn("rate limiter unavailable, failing open", "bucket", bucket, "err", err)
		return true
	}
	return n == 1
}

var distinctTargetsScript = redis.NewScript(`
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])
local member = ARGV[4]
redis.call('ZREMRANGEBYSCORE', key, 0, now - window)
redis.call('ZADD', key, now, member)
local count = redis.call('ZCARD', key)
redis.call('EXPIRE', key, window + 5)
if count > limit then return 0 end
return 1
`)

// AllowDistinctTargets throttles fan-out: a compromised account blasting many
// recipients inside one window is the classic spam / harassment signature.
func (l *Limiter) AllowDistinctTargets(ctx context.Context, subject, target string, limit int, window time.Duration) bool {
	if limit <= 0 {
		return true
	}
	key := "paperplane:fanout:" + subject
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	n, err := distinctTargetsScript.Run(ctx, l.rdb, []string{key},
		time.Now().Unix(), int(window.Seconds()), limit, target).Int()
	if err != nil {
		l.log.Warn("fan-out limiter unavailable, failing open", "err", err)
		return true
	}
	return n == 1
}

// IsBanned reads the ban ZSET written by the Kotlin policy sidecar.
func (l *Limiter) IsBanned(ctx context.Context, ip string) bool {
	if ip == "" {
		return false
	}
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	score, err := l.rdb.ZScore(ctx, banKey, ip).Result()
	if errors.Is(err, redis.Nil) {
		return false
	}
	if err != nil {
		l.log.Warn("ban lookup failed, failing open", "err", err)
		return false
	}
	return score > float64(time.Now().Unix())
}

// PublishEvent feeds the Kotlin sidecar through a capped Redis Stream.
// Only identifiers and coarse counters are published — never ciphertext.
func (l *Limiter) PublishEvent(ctx context.Context, kind string, fields map[string]string) {
	if l.stream == "" {
		return
	}
	values := map[string]any{
		"kind": kind,
		"ts":   strconv.FormatInt(time.Now().Unix(), 10),
	}
	for k, v := range fields {
		values[k] = v
	}
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	err := l.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: l.stream,
		MaxLen: 20000,
		Approx: true,
		Values: values,
	}).Err()
	if err != nil {
		l.log.Warn("event stream unavailable", "kind", kind, "err", err)
	}
}

// CountReconnect increments a short-lived counter used by the sidecar's
// reconnect-velocity rule.
func (l *Limiter) CountReconnect(ctx context.Context, accountID string) int64 {
	key := "paperplane:reconnect:" + accountID
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	n, err := l.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0
	}
	_ = l.rdb.Expire(ctx, key, 30*time.Second).Err()
	return n
}
```

### `server/internal/httpapi/handlers_auth.go`

```go
package httpapi

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/paperplane/paperplane/server/internal/auth"
	"github.com/paperplane/paperplane/server/internal/store"
)

const (
	nonceTTL    = 2 * time.Minute
	ticketTTL   = 30 * time.Second
	sealedTTL   = 24 * time.Hour
	sealedBatch = 5
)

type challengeRequest struct {
	Username string `json:"username"`
	Intent   string `json:"intent"` // "register" | "login"
}

type challengeResponse struct {
	AccountID string `json:"account_id"`
	Nonce     string `json:"nonce"`
	ExpiresIn int    `json:"expires_in"`
}

// handleChallenge issues a single-use nonce bound to the derived account id.
// For login we deliberately answer identically whether or not the account
// exists, so the endpoint cannot be used as a username oracle.
func (s *Server) handleChallenge(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "challenge", ip, s.cfg.Defense.RateLimits.ChallengePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many challenges")
		return
	}
	var req challengeRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	username := auth.NormalizeUsername(req.Username)
	if !auth.ValidUsername(username) {
		writeErr(w, http.StatusBadRequest, "invalid_username", "username must be 3-32 chars: a-z 0-9 . _ -")
		return
	}
	accountID := auth.AccountID(s.cfg.Crypto.Pepper, username)

	ctx, cancel := timeout(r, 3*time.Second)
	defer cancel()

	switch strings.ToLower(strings.TrimSpace(req.Intent)) {
	case "register":
		exists, err := s.db.AccountExists(ctx, accountID)
		if err != nil {
			s.log.Error("account lookup failed", "err", err)
			writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
			return
		}
		if exists {
			writeErr(w, http.StatusConflict, "account_exists", "username already registered")
			return
		}
	case "login", "":
	default:
		writeErr(w, http.StatusBadRequest, "invalid_intent", "intent must be login or register")
		return
	}

	nonce, err := auth.NewNonce()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "rng_unavailable", "could not issue nonce")
		return
	}
	if err := s.rdb.Set(ctx, nonceKey(accountID, nonce), "1", nonceTTL).Err(); err != nil {
		s.log.Error("nonce store failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, challengeResponse{
		AccountID: accountID, Nonce: nonce, ExpiresIn: int(nonceTTL.Seconds()),
	})
}

type registerRequest struct {
	Username     string `json:"username"`
	Argon2IDHash string `json:"argon2id_hash"`
	Nonce        string `json:"nonce"`
	Proof        string `json:"proof"`
}

type sessionResponse struct {
	Token     string `json:"token"`
	AccountID string `json:"account_id"`
	ExpiresIn int    `json:"expires_in"`
}

// handleRegister stores only the PHC string. The client-proved nonce ensures
// the caller actually holds the credential it is asking us to bind.
func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "register", ip, s.cfg.Defense.RateLimits.ChallengePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many registrations")
		return
	}
	var req registerRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	username := auth.NormalizeUsername(req.Username)
	if !auth.ValidUsername(username) {
		writeErr(w, http.StatusBadRequest, "invalid_username", "invalid username")
		return
	}
	accountID := auth.AccountID(s.cfg.Crypto.Pepper, username)

	mem, iter, par, err := auth.ParsePHC(req.Argon2IDHash)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_credential", err.Error())
		return
	}
	// Refuse downgraded parameters: a client must not register with cheaper
	// Argon2 settings than the server advertises.
	if mem < s.cfg.Crypto.Argon2.MemoryKiB || iter < s.cfg.Crypto.Argon2.Iterations ||
		par < s.cfg.Crypto.Argon2.Parallelism {
		writeErr(w, http.StatusBadRequest, "weak_credential", "argon2id parameters below server policy")
		return
	}

	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	if !s.consumeNonce(ctx, accountID, req.Nonce) {
		writeErr(w, http.StatusBadRequest, "invalid_nonce", "nonce missing, reused or expired")
		return
	}
	if !auth.VerifyProof(req.Argon2IDHash, req.Nonce, req.Proof) {
		writeErr(w, http.StatusUnauthorized, "invalid_proof", "proof does not match credential")
		return
	}

	if err := s.db.CreateAccount(ctx, accountID, req.Argon2IDHash); err != nil {
		if errors.Is(err, store.ErrAccountExists) {
			writeErr(w, http.StatusConflict, "account_exists", "username already registered")
			return
		}
		s.log.Error("create account failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}

	token, err := s.jwt.Issue(accountID)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "token_error", "could not issue token")
		return
	}
	s.limiter.PublishEvent(r.Context(), "register", map[string]string{"account": accountID, "ip": ip})
	writeJSON(w, http.StatusCreated, sessionResponse{
		Token: token, AccountID: accountID, ExpiresIn: int(s.jwt.TTL().Seconds()),
	})
}

type loginRequest struct {
	Username string `json:"username"`
	Nonce    string `json:"nonce"`
	Proof    string `json:"proof"`
}

// handleLogin never receives the credential: the client proves knowledge of
// the stored hash by HMAC-ing the nonce. A wrong password simply fails the
// constant-time proof check.
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "login", ip, s.cfg.Defense.RateLimits.ChallengePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many attempts")
		return
	}
	var req loginRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	username := auth.NormalizeUsername(req.Username)
	if !auth.ValidUsername(username) {
		writeErr(w, http.StatusUnauthorized, "invalid_credentials", "invalid credentials")
		return
	}
	accountID := auth.AccountID(s.cfg.Crypto.Pepper, username)

	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	used := s.consumeNonce(ctx, accountID, req.Nonce)
	stored, err := s.db.AuthHash(ctx, accountID)
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		s.log.Error("auth hash lookup failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	// Burn the nonce even when the account is unknown, and keep the response
	// identical for both failure paths.
	if !used || errors.Is(err, store.ErrNotFound) ||
		!auth.VerifyProof(stored, req.Nonce, req.Proof) {
		writeErr(w, http.StatusUnauthorized, "invalid_credentials", "invalid credentials")
		return
	}

	token, err := s.jwt.Issue(accountID)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "token_error", "could not issue token")
		return
	}
	s.db.TouchAccount(ctx, accountID)
	writeJSON(w, http.StatusOK, sessionResponse{
		Token: token, AccountID: accountID, ExpiresIn: int(s.jwt.TTL().Seconds()),
	})
}

func (s *Server) consumeNonce(ctx context.Context, accountID, nonce string) bool {
	if nonce == "" || len(nonce) > 64 {
		return false
	}
	val, err := s.rdb.GetDel(ctx, nonceKey(accountID, nonce)).Result()
	if errors.Is(err, redis.Nil) {
		return false
	}
	return err == nil && val != ""
}

func nonceKey(accountID, nonce string) string {
	return "paperplane:nonce:" + accountID + ":" + nonce
}

// handleWSTicket trades a long-lived JWT for a 30-second single-use ticket.
// Tokens therefore never appear in URLs or proxy access logs.
func (s *Server) handleWSTicket(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	ticket, err := auth.NewTicket()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "rng_unavailable", "could not issue ticket")
		return
	}
	ctx, cancel := timeout(r, 3*time.Second)
	defer cancel()
	if err := s.rdb.Set(ctx, "paperplane:wsticket:"+ticket, accountID, ticketTTL).Err(); err != nil {
		s.log.Error("ticket store failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ticket": ticket, "expires_in": int(ticketTTL.Seconds())})
}

// handleSealedTokens mints blind sender tokens. Each token can be attached to
// exactly one envelope, after which the server stores no sender at all.
func (s *Server) handleSealedTokens(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	ctx, cancel := timeout(r, 3*time.Second)
	defer cancel()

	tokens := make([]string, 0, sealedBatch)
	for i := 0; i < sealedBatch; i++ {
		t, err := auth.NewTicket()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, "rng_unavailable", "could not mint token")
			return
		}
		if err := s.rdb.Set(ctx, "paperplane:sealed:"+t, accountID, sealedTTL).Err(); err != nil {
			s.log.Error("sealed token store failed", "err", err)
			writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
			return
		}
		tokens = append(tokens, t)
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"tokens": tokens, "expires_in": int(sealedTTL.Seconds()),
	})
}
```

### `server/internal/httpapi/handlers_envelopes.go`

```go
package httpapi

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/gorilla/websocket"

	"github.com/paperplane/paperplane/server/internal/store"
	"github.com/paperplane/paperplane/server/internal/ws"
)

const maxEnvelopeBytes = 1 << 20 // decoded ciphertext cap

type envelopeRequest struct {
	To           string `json:"to"`
	ClientMsgID  string `json:"client_msg_id"`
	CipherType   int    `json:"cipher_type"`
	Body         string `json:"body"`
	SealedSender bool   `json:"sealed_sender"`
	SenderToken  string `json:"sender_token"`
}

type envelopeOut struct {
	ID          int64     `json:"id"`
	From        string    `json:"from,omitempty"`
	CipherType  int       `json:"cipher_type"`
	Body        string    `json:"body"`
	ClientMsgID string    `json:"client_msg_id"`
	SentAt      time.Time `json:"sent_at"`
}

func toOut(e store.Envelope) envelopeOut {
	return envelopeOut{
		ID: e.ID, From: e.SenderID, CipherType: e.CipherType,
		Body: base64.StdEncoding.EncodeToString(e.Body),
		ClientMsgID: e.ClientMsgID, SentAt: e.SentAt,
	}
}

func envelopeFrame(e store.Envelope) ws.EnvelopeFrame {
	return ws.EnvelopeFrame{
		Type: "envelope", ID: e.ID, From: e.SenderID, CipherType: e.CipherType,
		Body: base64.StdEncoding.EncodeToString(e.Body),
		ClientMsgID: e.ClientMsgID, SentAt: e.SentAt,
	}
}

// handlePostEnvelope accepts an opaque ciphertext. The server cannot decrypt
// it, does not log it, and deletes it once the recipient acknowledges delivery.
func (s *Server) handlePostEnvelope(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	ip := s.clientIP(r)

	if !s.limiter.AllowTokenBucket(r.Context(), "envelope", accountID, s.cfg.Defense.RateLimits.EnvelopePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "envelope rate exceeded")
		return
	}

	var req envelopeRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	if !accountIDPattern.MatchString(req.To) {
		writeErr(w, http.StatusBadRequest, "invalid_recipient", "recipient must be a 64-char account id")
		return
	}
	if req.To == accountID {
		writeErr(w, http.StatusBadRequest, "invalid_recipient", "cannot send to self")
		return
	}
	if len(req.ClientMsgID) == 0 || len(req.ClientMsgID) > 80 {
		writeErr(w, http.StatusBadRequest, "invalid_client_msg_id", "client_msg_id must be 1-80 chars")
		return
	}
	if req.CipherType != 3 && req.CipherType != 4 {
		// 3 = SignalMessage, 4 = PreKeySignalMessage (libsignal wire types).
		writeErr(w, http.StatusBadRequest, "invalid_cipher_type", "cipher_type must be 3 or 4")
		return
	}
	body, err := base64.StdEncoding.DecodeString(req.Body)
	if err != nil {
		body, err = base64.RawStdEncoding.DecodeString(req.Body)
	}
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body_encoding", "ciphertext must be base64")
		return
	}
	if len(body) == 0 || len(body) > maxEnvelopeBytes {
		writeErr(w, http.StatusRequestEntityTooLarge, "envelope_too_large", "ciphertext must be 1 byte - 1 MiB")
		return
	}

	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	// Fan-out guard: one sender talking to too many distinct recipients inside
	// a window is blocked before any database write happens.
	if !s.limiter.AllowDistinctTargets(ctx, accountID, req.To,
		s.cfg.Defense.RateLimits.EnvelopeDistinctTargets, s.cfg.EnvelopeWindow()) {
		writeErr(w, http.StatusTooManyRequests, "fanout_limited", "too many distinct recipients")
		return
	}

	sender := accountID
	if req.SealedSender {
		if req.SenderToken == "" {
			writeErr(w, http.StatusBadRequest, "missing_sender_token", "sealed sender requires a token")
			return
		}
		owner, err := s.rdb.GetDel(ctx, "paperplane:sealed:"+req.SenderToken).Result()
		if errors.Is(err, redis.Nil) || (err == nil && owner != accountID) {
			writeErr(w, http.StatusUnauthorized, "invalid_sender_token", "sealed sender token rejected")
			return
		}
		if err != nil && !errors.Is(err, redis.Nil) {
			s.log.Error("sealed token lookup failed", "err", err)
			writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
			return
		}
		sender = "" // route metadata intentionally not persisted
	}

	exists, err := s.db.AccountExists(ctx, req.To)
	if err != nil {
		s.log.Error("recipient lookup failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	if !exists {
		writeErr(w, http.StatusNotFound, "no_recipient", "recipient has no account")
		return
	}

	stored := store.Envelope{
		TargetID: req.To, SenderID: sender, CipherType: req.CipherType,
		Body: body, ClientMsgID: req.ClientMsgID,
	}
	id, err := s.db.InsertEnvelope(ctx, stored, s.cfg.Database.EnvelopeTTLHour)
	if err != nil {
		s.log.Error("envelope insert failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	stored.ID = id
	stored.SentAt = time.Now().UTC()

	payload, _ := jsonMarshal(envelopeFrame(stored))
	delivered := s.hub.Deliver(req.To, payload)
	if delivered {
		// Once handed to a live socket, drop sender metadata immediately.
		s.db.StripSender(ctx, []int64{id})
	}
	s.limiter.PublishEvent(ctx, "envelope", map[string]string{
		"account": accountID, "ip": ip, "target": req.To,
		"sealed": strconv.FormatBool(req.SealedSender),
	})

	writeJSON(w, http.StatusCreated, map[string]any{
		"id": id, "delivered": delivered,
	})
}

func (s *Server) handlePullEnvelopes(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	limit := 50
	if raw := r.URL.Query().Get("limit"); raw != "" {
		if n, err := strconv.Atoi(raw); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	list, err := s.db.PendingEnvelopes(ctx, accountID, limit)
	if err != nil {
		s.log.Error("envelope pull failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	out := make([]envelopeOut, 0, len(list))
	for _, e := range list {
		out = append(out, toOut(e))
	}
	writeJSON(w, http.StatusOK, map[string]any{"envelopes": out, "count": len(out)})
}

type ackRequest struct {
	IDs []int64 `json:"ids"`
}

func (s *Server) handleAckEnvelopes(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	var req ackRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	if len(req.IDs) == 0 || len(req.IDs) > 200 {
		writeErr(w, http.StatusBadRequest, "invalid_ids", "ids must contain 1-200 entries")
		return
	}
	for _, id := range req.IDs {
		if id <= 0 {
			writeErr(w, http.StatusBadRequest, "invalid_ids", "ids must be positive")
			return
		}
	}
	ctx, cancel := timeout(r, 5*time.Second)
	defer cancel()
	n, err := s.db.AckEnvelopes(ctx, accountID, req.IDs)
	if err != nil {
		s.log.Error("envelope ack failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"acked": n})
}

// ------------------------------------------------------------- websocket

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.originAllowed(r.Header.Get("Origin")) {
		writeErr(w, http.StatusForbidden, "origin_denied", "websocket origin not allowed")
		return
	}
	if s.limiter.IsBanned(r.Context(), ip) {
		writeErr(w, http.StatusForbidden, "banned", "temporarily blocked")
		return
	}
	if s.hub.ConnectionsForIP(ip) >= s.cfg.Defense.RateLimits.WSPerIP {
		writeErr(w, http.StatusTooManyRequests, "ws_limit", "too many concurrent sockets from this address")
		return
	}

	ticket := r.URL.Query().Get("ticket")
	if ticket == "" || len(ticket) > 128 {
		writeErr(w, http.StatusUnauthorized, "missing_ticket", "single-use ticket required")
		return
	}
	ctx, cancel := timeout(r, 5*time.Second)
	accountID, err := s.rdb.GetDel(ctx, "paperplane:wsticket:"+ticket).Result()
	cancel()
	if errors.Is(err, redis.Nil) {
		writeErr(w, http.StatusUnauthorized, "invalid_ticket", "ticket expired or already used")
		return
	}
	if err != nil {
		s.log.Error("ticket lookup failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	if !accountIDPattern.MatchString(accountID) {
		writeErr(w, http.StatusUnauthorized, "invalid_ticket", "ticket rejected")
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Exact whitelist only: never `return true`, never compare r.Host
		// (Caddy rewrites Host but preserves Origin).
		CheckOrigin: func(req *http.Request) bool {
			return s.originAllowed(req.Header.Get("Origin"))
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Warn("websocket upgrade failed", "err", err)
		return
	}

	client := s.hub.Add(accountID, ip, conn)
	client.Start(r.Context())
	defer client.Close()

	conn.SetReadLimit(64 << 10)
	_ = conn.SetReadDeadline(time.Now().Add(90 * time.Second))
	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(90 * time.Second))
	})

	ready, _ := jsonMarshal(ws.ReadyFrame{
		Type: "ready", AccountID: accountID, ServerTS: time.Now().Unix(),
	})
	client.Send(ready)

	// Catch-up: deliver everything queued while the device was offline.
	pullCtx, pullCancel := timeout(r, 8*time.Second)
	pending, err := s.db.PendingEnvelopes(pullCtx, accountID, 50)
	pullCancel()
	if err != nil {
		s.log.Warn("pending pull on connect failed", "err", err)
	}
	pushed := make([]int64, 0, len(pending))
	for _, e := range pending {
		payload, _ := jsonMarshal(envelopeFrame(e))
		if client.Send(payload) {
			pushed = append(pushed, e.ID)
		}
	}
	if len(pushed) > 0 {
		stripCtx, stripCancel := timeout(r, 3*time.Second)
		s.db.StripSender(stripCtx, pushed)
		stripCancel()
	}

	s.limiter.PublishEvent(r.Context(), "connect", map[string]string{
		"account": accountID, "ip": ip,
		"reconnects": strconv.FormatInt(s.limiter.CountReconnect(r.Context(), accountID), 10),
	})

	var inbound struct {
		Type string  `json:"type"`
		IDs  []int64 `json:"ids"`
	}
	for {
		if err := conn.ReadJSON(&inbound); err != nil {
			return
		}
		switch inbound.Type {
		case "ack":
			if len(inbound.IDs) == 0 || len(inbound.IDs) > 200 {
				continue
			}
			ackCtx, ackCancel := timeout(r, 5*time.Second)
			if _, err := s.db.AckEnvelopes(ackCtx, accountID, inbound.IDs); err != nil {
				s.log.Warn("ws ack failed", "err", err)
			}
			ackCancel()
		case "ping":
			client.Send([]byte(`{"type":"pong"}`))
		}
	}
}
```

### `server/internal/httpapi/handlers_keys.go`

```go
package httpapi

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/paperplane/paperplane/server/internal/store"
)

var accountIDPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

type preKeyJSON struct {
	ID        int64  `json:"id"`
	Public    string `json:"public"`
	Signature string `json:"signature,omitempty"`
}

type oneTimePreKeyJSON struct {
	ID     int64  `json:"id"`
	Public string `json:"public"`
}

type bundleRequest struct {
	IdentityKey    string             `json:"identity_key"`
	SignedPreKey   preKeyJSON         `json:"signed_prekey"`
	KyberPreKey    *preKeyJSON        `json:"kyber_prekey,omitempty"`
	OneTimePreKeys []oneTimePreKeyJSON `json:"one_time_prekeys"`
}

type bundleResponse struct {
	AccountID     string              `json:"account_id"`
	IdentityKey   string              `json:"identity_key"`
	SignedPreKey  preKeyJSON          `json:"signed_prekey"`
	KyberPreKey   *preKeyJSON         `json:"kyber_prekey,omitempty"`
	OneTimePreKey *oneTimePreKeyJSON  `json:"one_time_prekey,omitempty"`
}

func (s *Server) handleUploadKeys(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "prekey", accountID, s.cfg.Defense.RateLimits.PreKeyPerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many pre-key updates")
		return
	}
	var req bundleRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}

	identity, err := decodeKey(req.IdentityKey, 32, 128)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_identity_key", err.Error())
		return
	}
	signed, err := decodePreKey(req.SignedPreKey, true)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_signed_prekey", err.Error())
		return
	}
	if req.KyberPreKey != nil && req.KyberPreKey.Public != "" {
		if _, err := decodePreKey(*req.KyberPreKey, false); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid_kyber_prekey", err.Error())
			return
		}
	}
	if len(req.OneTimePreKeys) == 0 || len(req.OneTimePreKeys) > 200 {
		writeErr(w, http.StatusBadRequest, "invalid_one_time_prekeys", "must upload 1-200 one-time prekeys")
		return
	}

	otps := make([]store.OneTimePreKey, 0, len(req.OneTimePreKeys))
	for _, k := range req.OneTimePreKeys {
		if k.ID <= 0 {
			writeErr(w, http.StatusBadRequest, "invalid_one_time_prekeys", "prekey id must be positive")
			return
		}
		pub, err := decodeKey(k.Public, 32, 128)
		if err != nil {
			writeErr(w, http.StatusBadRequest, "invalid_one_time_prekeys", err.Error())
			return
		}
		otps = append(otps, store.OneTimePreKey{ID: k.ID, PublicKey: pub})
	}

	bundle := store.Bundle{
		AccountID:             accountID,
		IdentityKey:           identity,
		SignedPreKeyID:        signed.ID,
		SignedPreKeyPublic:    signed.PublicKey,
		SignedPreKeySignature: signed.Signature,
	}
	if req.KyberPreKey != nil && req.KyberPreKey.Public != "" {
		kyber, _ := decodePreKey(*req.KyberPreKey, false)
		bundle.KyberPreKeyID = kyber.ID
		bundle.KyberPreKeyPublic = kyber.PublicKey
		bundle.KyberPreKeySignature = kyber.Signature
	}

	ctx, cancel := timeout(r, 5*time.Second)
	defer cancel()
	if err := s.db.UpsertBundle(ctx, bundle, otps); err != nil {
		s.log.Error("bundle upload failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"one_time_prekeys_stored": len(otps),
		"identity_key":            req.IdentityKey,
	})
}

func (s *Server) handleGetBundle(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "prekey", accountID, s.cfg.Defense.RateLimits.PreKeyPerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many pre-key fetches")
		return
	}
	target := r.PathValue("accountID")
	if !accountIDPattern.MatchString(target) {
		writeErr(w, http.StatusBadRequest, "invalid_account", "account id must be 64 lowercase hex chars")
		return
	}
	ctx, cancel := timeout(r, 5*time.Second)
	defer cancel()

	b, err := s.db.TakeBundle(ctx, target)
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, http.StatusNotFound, "no_bundle", "recipient has not published keys")
		return
	}
	if err != nil {
		s.log.Error("bundle fetch failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}

	resp := bundleResponse{
		AccountID:   b.AccountID,
		IdentityKey: b64(b.IdentityKey),
		SignedPreKey: preKeyJSON{
			ID: b.SignedPreKeyID, Public: b64(b.SignedPreKeyPublic),
			Signature: b64(b.SignedPreKeySignature),
		},
	}
	if len(b.KyberPreKeyPublic) > 0 {
		resp.KyberPreKey = &preKeyJSON{
			ID: b.KyberPreKeyID, Public: b64(b.KyberPreKeyPublic),
			Signature: b64(b.KyberPreKeySignature),
		}
	}
	if b.OneTimePreKey != nil {
		resp.OneTimePreKey = &oneTimePreKeyJSON{
			ID: b.OneTimePreKey.ID, Public: b64(b.OneTimePreKey.PublicKey),
		}
	}
	writeJSON(w, http.StatusOK, resp)
}

// ------------------------------------------------------------------ decoding

func b64(b []byte) string { return base64.StdEncoding.EncodeToString(b) }

func decodeKey(s string, minLen, maxLen int) ([]byte, error) {
	raw, err := decodeB64(s)
	if err != nil {
		return nil, errors.New("value is not valid base64")
	}
	if len(raw) < minLen || len(raw) > maxLen {
		return nil, fmt.Errorf("key length %d outside %d-%d bytes", len(raw), minLen, maxLen)
	}
	return raw, nil
}

func decodePreKey(k preKeyJSON, requireSignature bool) (struct {
	ID        int64
	PublicKey []byte
	Signature []byte
}, error) {
	var out struct {
		ID        int64
		PublicKey []byte
		Signature []byte
	}
	if k.ID <= 0 {
		return out, errors.New("prekey id must be positive")
	}
	pub, err := decodeKey(k.Public, 32, 128)
	if err != nil {
		return out, err
	}
	out.ID, out.PublicKey = k.ID, pub
	if k.Signature != "" {
		sig, err := decodeKey(k.Signature, 32, 128)
		if err != nil {
			return out, err
		}
		out.Signature = sig
	} else if requireSignature {
		return out, errors.New("signed prekey requires a signature")
	}
	return out, nil
}

func decodeB64(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("empty")
	}
	for _, enc := range []*base64.Encoding{
		base64.StdEncoding, base64.RawStdEncoding, base64.URLEncoding, base64.RawURLEncoding,
	} {
		if raw, err := enc.DecodeString(s); err == nil {
			return raw, nil
		}
	}
	return nil, errors.New("invalid base64")
}
```

### `server/internal/httpapi/middleware.go`

```go
package httpapi

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/paperplane/paperplane/server/internal/auth"
)

type contextKey string

const accountKey contextKey = "paperplane.account_id"

func accountFrom(r *http.Request) string {
	v, _ := r.Context().Value(accountKey).(string)
	return v
}

func timeout(r *http.Request, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), d)
}

// common wraps the mux with the middleware stack. Order matters: recover first
// so a panic in any later layer still yields a 500 instead of a dead socket.
func (s *Server) common(next http.Handler) http.Handler {
	return s.recoverer(s.securityHeaders(s.accessLog(s.cors(next))))
}

func (s *Server) recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				s.log.Error("panic recovered", "path", r.URL.Path, "panic", rec)
				writeErr(w, http.StatusInternalServerError, "internal_error", "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (s *Server) securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Referrer-Policy", "no-referrer")
		h.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
		h.Set("Permissions-Policy", "interest-cohort=()")
		if s.cfg.IsProduction() {
			h.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		}
		next.ServeHTTP(w, r)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

func (s *Server) accessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		// Only method/path/status/ip/duration — never headers, tokens or bodies.
		s.log.Info("http",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"ip", s.clientIP(r),
			"dur_ms", time.Since(start).Milliseconds(),
		)
	})
}

func (s *Server) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && s.originAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Max-Age", "600")
		}
		if r.Method == http.MethodOptions {
			if origin == "" || !s.originAllowed(origin) {
				writeErr(w, http.StatusForbidden, "origin_denied", "origin not allowed")
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// requireAuth validates the bearer token and injects account_id into context.
func (s *Server) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			writeErr(w, http.StatusUnauthorized, "missing_token", "authorization header required")
			return
		}
		accountID, err := s.jwt.Verify(strings.TrimSpace(strings.TrimPrefix(header, "Bearer ")))
		if err != nil {
			code := "invalid_token"
			if err == auth.ErrExpiredToken {
				code = "token_expired"
			}
			writeErr(w, http.StatusUnauthorized, code, "token rejected")
			return
		}
		next(w, r.WithContext(context.WithValue(r.Context(), accountKey, accountID)))
	}
}
```

### `server/internal/httpapi/router.go`

```go
// Package httpapi is the only HTTP surface of the server. Every handler is
// deliberately small: validate, rate-limit, touch the store, return JSON.
// Ciphertext is opaque here — no handler inspects a message body.
package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/paperplane/paperplane/server/internal/auth"
	"github.com/paperplane/paperplane/server/internal/config"
	"github.com/paperplane/paperplane/server/internal/defense"
	"github.com/paperplane/paperplane/server/internal/store"
	"github.com/paperplane/paperplane/server/internal/ws"
)

const maxBodyBytes = 2 << 20 // 2 MiB encoded; decoded ciphertext capped at 1 MiB

type Server struct {
	cfg     *config.Config
	db      *store.Postgres
	rdb     *redis.Client
	limiter *defense.Limiter
	hub     *ws.Hub
	jwt     *auth.Issuer
	log     *slog.Logger
	started time.Time
}

func New(cfg *config.Config, db *store.Postgres, rdb *redis.Client, lim *defense.Limiter,
	hub *ws.Hub, jwt *auth.Issuer, log *slog.Logger) *Server {
	return &Server{
		cfg: cfg, db: db, rdb: rdb, limiter: lim, hub: hub, jwt: jwt,
		log: log, started: time.Now(),
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("GET /api/v1/params", s.handleParams)

	mux.HandleFunc("POST /api/v1/auth/challenge", s.handleChallenge)
	mux.HandleFunc("POST /api/v1/auth/register", s.handleRegister)
	mux.HandleFunc("POST /api/v1/auth/login", s.handleLogin)

	mux.HandleFunc("POST /api/v1/ws-ticket", s.requireAuth(s.handleWSTicket))
	mux.HandleFunc("POST /api/v1/sealed-tokens", s.requireAuth(s.handleSealedTokens))

	mux.HandleFunc("POST /api/v1/keys", s.requireAuth(s.handleUploadKeys))
	mux.HandleFunc("GET /api/v1/keys/{accountID}", s.requireAuth(s.handleGetBundle))

	mux.HandleFunc("POST /api/v1/envelopes", s.requireAuth(s.handlePostEnvelope))
	mux.HandleFunc("GET /api/v1/envelopes", s.requireAuth(s.handlePullEnvelopes))
	mux.HandleFunc("POST /api/v1/envelopes/ack", s.requireAuth(s.handleAckEnvelopes))

	mux.HandleFunc("GET /ws", s.handleWS)
	mux.HandleFunc("/", s.handleNotFound)

	return s.common(mux)
}

// ------------------------------------------------------------------ helpers

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	if v == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Debug("write response", "err", err)
	}
}

func writeErr(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": map[string]string{"code": code, "message": message},
	})
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

func (s *Server) clientIP(r *http.Request) string {
	// The service binds loopback and is reachable only through Caddy, which
	// sets X-Forwarded-For. Trust the left-most well-formed address.
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		first := strings.TrimSpace(strings.Split(xff, ",")[0])
		if ip := net.ParseIP(first); ip != nil {
			return first
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func (s *Server) originAllowed(raw string) bool {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		// Native clients send no Origin; browsers always do.
		return true
	}
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	normalized := strings.ToLower(u.Scheme + "://" + u.Host)
	for _, allowed := range s.cfg.Server.WSAllowedOrigins {
		au, err := url.Parse(strings.TrimSpace(allowed))
		if err != nil || au.Scheme == "" || au.Host == "" {
			continue
		}
		if strings.ToLower(au.Scheme+"://"+au.Host) == normalized {
			return true
		}
	}
	return false
}

func jsonMarshal(v any) ([]byte, error) { return json.Marshal(v) }

func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	writeErr(w, http.StatusNotFound, "not_found", "unknown route")
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := timeout(r, 2*time.Second)
	defer cancel()
	status := "ok"
	if err := s.db.Ping(ctx); err != nil {
		status = "degraded"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"status":   status,
		"uptime_s": int(time.Since(s.started).Seconds()),
		"time":     time.Now().UTC().Format(time.RFC3339),
	})
}

type paramsResponse struct {
	Project        string `json:"project"`
	PublicURL      string `json:"public_url"`
	Pepper         string `json:"pepper"`
	Argon2Memory   uint32 `json:"argon2_memory_kib"`
	Argon2Iter     uint32 `json:"argon2_iterations"`
	Argon2Parallel uint8  `json:"argon2_parallelism"`
	WSPath         string `json:"ws_path"`
	EnvelopeTTL    int    `json:"envelope_ttl_hours"`
	MaxEnvelope    int    `json:"max_envelope_bytes"`
	TokenTTL       int    `json:"token_ttl_seconds"`
}

// handleParams publishes exactly what a client needs to derive credentials.
// The pepper is a public domain separator by design: account ids stay
// irreversible because the server never stores a username.
func (s *Server) handleParams(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, paramsResponse{
		Project:        s.cfg.Project,
		PublicURL:      s.cfg.Server.PublicURL,
		Pepper:         s.cfg.Crypto.Pepper,
		Argon2Memory:   s.cfg.Crypto.Argon2.MemoryKiB,
		Argon2Iter:     s.cfg.Crypto.Argon2.Iterations,
		Argon2Parallel: s.cfg.Crypto.Argon2.Parallelism,
		WSPath:         "/ws",
		EnvelopeTTL:    s.cfg.Database.EnvelopeTTLHour,
		MaxEnvelope:    1 << 20,
		TokenTTL:       int(s.jwt.TTL().Seconds()),
	})
}

// unwrapBody normalises EOF (empty body) into a friendlier error.
func unwrapBody(err error) error {
	if errors.Is(err, io.EOF) {
		return errors.New("empty request body")
	}
	return err
}
```

### `server/internal/store/postgres.go`

```go
// Package store contains the durable state layer. PostgreSQL holds accounts,
// pre-key bundles and undelivered ciphertext envelopes; nothing here ever
// stores a plaintext message, a password or a reversible username.
package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAccountExists = errors.New("account already exists")
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse postgres dsn: %w", err)
	}
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	return &Postgres{pool: pool}, nil
}

func (p *Postgres) Ping(ctx context.Context) error { return p.pool.Ping(ctx) }
func (p *Postgres) Close()                         { p.pool.Close() }

// Migrate applies idempotent DDL. Statements are split on ';' because the
// extended protocol used by pgx accepts only one statement per Exec.
func (p *Postgres) Migrate(ctx context.Context, script string) error {
	for _, stmt := range strings.Split(script, ";") {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" || strings.HasPrefix(stmt, "--") && !strings.Contains(stmt, "\n") {
			continue
		}
		if _, err := p.pool.Exec(ctx, stmt); err != nil {
			return fmt.Errorf("migrate statement %.60q: %w", stmt, err)
		}
	}
	return nil
}

// ---------------------------------------------------------------- accounts

func (p *Postgres) CreateAccount(ctx context.Context, accountID, authHash string) error {
	tag, err := p.pool.Exec(ctx,
		`INSERT INTO accounts (account_id, auth_hash) VALUES ($1, $2)
		 ON CONFLICT (account_id) DO NOTHING`, accountID, authHash)
	if err != nil {
		return fmt.Errorf("insert account: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrAccountExists
	}
	return nil
}

func (p *Postgres) AuthHash(ctx context.Context, accountID string) (string, error) {
	var hash string
	err := p.pool.QueryRow(ctx, `SELECT auth_hash FROM accounts WHERE account_id = $1`, accountID).Scan(&hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("select auth hash: %w", err)
	}
	return hash, nil
}

func (p *Postgres) AccountExists(ctx context.Context, accountID string) (bool, error) {
	var exists bool
	err := p.pool.QueryRow(ctx,
		`SELECT EXISTS (SELECT 1 FROM accounts WHERE account_id = $1)`, accountID).Scan(&exists)
	return exists, err
}

func (p *Postgres) TouchAccount(ctx context.Context, accountID string) {
	_, _ = p.pool.Exec(ctx, `UPDATE accounts SET last_seen_at = now() WHERE account_id = $1`, accountID)
}

// ---------------------------------------------------------------- prekeys

type OneTimePreKey struct {
	ID        int64
	PublicKey []byte
}

// Bundle is the public half of a Signal identity. Kyber fields are optional
// because pure-Dart libsignal implements X3DH; the wire format reserves them so
// a PQXDH-capable client can populate them without a schema migration.
type Bundle struct {
	AccountID             string
	IdentityKey           []byte
	SignedPreKeyID        int64
	SignedPreKeyPublic    []byte
	SignedPreKeySignature []byte
	KyberPreKeyID         int64
	KyberPreKeyPublic     []byte
	KyberPreKeySignature  []byte
	OneTimePreKey         *OneTimePreKey
}

func (p *Postgres) UpsertBundle(ctx context.Context, b Bundle, oneTime []OneTimePreKey) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	_, err = tx.Exec(ctx, `
		INSERT INTO prekey_bundles (
			account_id, identity_key, signed_prekey_id, signed_prekey_public,
			signed_prekey_signature, kyber_prekey_id, kyber_prekey_public,
			kyber_prekey_signature, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8, now())
		ON CONFLICT (account_id) DO UPDATE SET
			identity_key = EXCLUDED.identity_key,
			signed_prekey_id = EXCLUDED.signed_prekey_id,
			signed_prekey_public = EXCLUDED.signed_prekey_public,
			signed_prekey_signature = EXCLUDED.signed_prekey_signature,
			kyber_prekey_id = EXCLUDED.kyber_prekey_id,
			kyber_prekey_public = EXCLUDED.kyber_prekey_public,
			kyber_prekey_signature = EXCLUDED.kyber_prekey_signature,
			updated_at = now()`,
		b.AccountID, b.IdentityKey, b.SignedPreKeyID, b.SignedPreKeyPublic,
		b.SignedPreKeySignature, b.KyberPreKeyID, b.KyberPreKeyPublic,
		b.KyberPreKeySignature)
	if err != nil {
		return fmt.Errorf("upsert bundle: %w", err)
	}

	for _, k := range oneTime {
		if _, err := tx.Exec(ctx,
			`INSERT INTO one_time_prekeys (account_id, prekey_id, public_key)
			 VALUES ($1,$2,$3) ON CONFLICT (account_id, prekey_id) DO NOTHING`,
			b.AccountID, k.ID, k.PublicKey); err != nil {
			return fmt.Errorf("insert one-time prekey: %w", err)
		}
	}
	return tx.Commit(ctx)
}

// TakeBundle returns the published bundle and atomically consumes exactly one
// one-time prekey (SKIP LOCKED keeps concurrent fetchers from colliding).
func (p *Postgres) TakeBundle(ctx context.Context, accountID string) (*Bundle, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	b := &Bundle{AccountID: accountID}
	err = tx.QueryRow(ctx, `
		SELECT identity_key,
		       signed_prekey_id, signed_prekey_public, signed_prekey_signature,
		       COALESCE(kyber_prekey_id, 0),
		       COALESCE(kyber_prekey_public, '\x'::bytea),
		       COALESCE(kyber_prekey_signature, '\x'::bytea)
		FROM prekey_bundles WHERE account_id = $1`, accountID).
		Scan(&b.IdentityKey, &b.SignedPreKeyID, &b.SignedPreKeyPublic,
			&b.SignedPreKeySignature, &b.KyberPreKeyID, &b.KyberPreKeyPublic,
			&b.KyberPreKeySignature)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("load bundle: %w", err)
	}

	var otpID int64
	var otpPublic []byte
	err = tx.QueryRow(ctx, `
		SELECT prekey_id, public_key FROM one_time_prekeys
		WHERE account_id = $1 ORDER BY prekey_id LIMIT 1 FOR UPDATE SKIP LOCKED`, accountID).
		Scan(&otpID, &otpPublic)
	switch {
	case err == nil:
		if _, delErr := tx.Exec(ctx,
			`DELETE FROM one_time_prekeys WHERE account_id = $1 AND prekey_id = $2`,
			accountID, otpID); delErr != nil {
			return nil, fmt.Errorf("consume one-time prekey: %w", delErr)
		}
		b.OneTimePreKey = &OneTimePreKey{ID: otpID, PublicKey: otpPublic}
	case errors.Is(err, pgx.ErrNoRows):
		// No OTP left is a normal state: X3DH can proceed without it.
	default:
		return nil, fmt.Errorf("take one-time prekey: %w", err)
	}
	return b, tx.Commit(ctx)
}

func (p *Postgres) CountOneTimePreKeys(ctx context.Context, accountID string) (int, error) {
	var n int
	err := p.pool.QueryRow(ctx,
		`SELECT count(*) FROM one_time_prekeys WHERE account_id = $1`, accountID).Scan(&n)
	return n, err
}

// ---------------------------------------------------------------- envelopes

type Envelope struct {
	ID          int64
	TargetID    string
	SenderID    string // empty when sealed sender
	CipherType  int
	Body        []byte
	ClientMsgID string
	SentAt      time.Time
}

func (p *Postgres) InsertEnvelope(ctx context.Context, e Envelope, ttlHours int) (int64, error) {
	var id int64
	var sender any
	if e.SenderID != "" {
		sender = e.SenderID
	}
	err := p.pool.QueryRow(ctx, `
		INSERT INTO envelopes (target_id, sender_id, cipher_type, body, client_msg_id, expires_at)
		VALUES ($1, $2, $3, $4, $5, now() + make_interval(hours => $6))
		RETURNING id`,
		e.TargetID, sender, e.CipherType, e.Body, e.ClientMsgID, ttlHours).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert envelope: %w", err)
	}
	return id, nil
}

func (p *Postgres) PendingEnvelopes(ctx context.Context, targetID string, limit int) ([]Envelope, error) {
	rows, err := p.pool.Query(ctx, `
		SELECT id, COALESCE(sender_id, ''), cipher_type, body, client_msg_id, sent_at
		FROM envelopes
		WHERE target_id = $1 AND expires_at > now()
		ORDER BY id LIMIT $2`, targetID, limit)
	if err != nil {
		return nil, fmt.Errorf("list envelopes: %w", err)
	}
	defer rows.Close()

	out := make([]Envelope, 0, limit)
	for rows.Next() {
		var e Envelope
		e.TargetID = targetID
		if err := rows.Scan(&e.ID, &e.SenderID, &e.CipherType, &e.Body, &e.ClientMsgID, &e.SentAt); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// StripSender drops routing metadata as soon as an envelope reaches a live
// socket. The ciphertext stays only until the recipient acknowledges it.
func (p *Postgres) StripSender(ctx context.Context, ids []int64) {
	if len(ids) == 0 {
		return
	}
	_, _ = p.pool.Exec(ctx, `UPDATE envelopes SET sender_id = NULL WHERE id = ANY($1::bigint[])`, ids)
}

func (p *Postgres) AckEnvelopes(ctx context.Context, targetID string, ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	tag, err := p.pool.Exec(ctx,
		`DELETE FROM envelopes WHERE target_id = $1 AND id = ANY($2::bigint[])`, targetID, ids)
	if err != nil {
		return 0, fmt.Errorf("ack envelopes: %w", err)
	}
	return tag.RowsAffected(), nil
}

func (p *Postgres) CleanupExpiredEnvelopes(ctx context.Context) (int64, error) {
	tag, err := p.pool.Exec(ctx, `DELETE FROM envelopes WHERE expires_at < now()`)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
```

### `server/internal/store/redis.go`

```go
package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/paperplane/paperplane/server/internal/config"
)

// NewRedis builds the shared Redis client. Redis holds only ephemeral state:
// nonces, WebSocket tickets, rate-limit counters, ban decisions and the
// defence event stream. No ciphertext and no credential is ever written here.
func NewRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr(),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: 4,
		DialTimeout:  3,
	})
}

func PingRedis(ctx context.Context, rdb *redis.Client) error {
	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping: %w", err)
	}
	return nil
}
```

### `server/internal/ws/hub.go`

```go
// Package ws owns the WebSocket fan-out. One account may have several live
// sockets (phone + laptop); each connection gets its own bounded send queue so
// a single slow reader can never block delivery to everyone else.
package ws

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 90 * time.Second
	pingPeriod     = 30 * time.Second
	maxMessageSize = 64 << 10
	sendQueueSize  = 32
)

type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	AccountID string
	IP        string

	send      chan []byte
	done      chan struct{}
	closeOnce sync.Once
}

type Hub struct {
	mu        sync.RWMutex
	byAccount map[string]map[*Client]struct{}
	byIP      map[string]int
	log       *slog.Logger
	closed    bool
}

func NewHub(log *slog.Logger) *Hub {
	return &Hub{
		byAccount: map[string]map[*Client]struct{}{},
		byIP:      map[string]int{},
		log:       log,
	}
}

// Add registers a freshly upgraded connection and returns the client handle.
func (h *Hub) Add(accountID, ip string, conn *websocket.Conn) *Client {
	c := &Client{
		hub:       h,
		conn:      conn,
		AccountID: accountID,
		IP:        ip,
		send:      make(chan []byte, sendQueueSize),
		done:      make(chan struct{}),
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.byAccount[accountID] == nil {
		h.byAccount[accountID] = map[*Client]struct{}{}
	}
	h.byAccount[accountID][c] = struct{}{}
	h.byIP[ip]++
	return c
}

func (h *Hub) remove(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if set, ok := h.byAccount[c.AccountID]; ok {
		delete(set, c)
		if len(set) == 0 {
			delete(h.byAccount, c.AccountID)
		}
	}
	if c.IP != "" {
		h.byIP[c.IP]--
		if h.byIP[c.IP] <= 0 {
			delete(h.byIP, c.IP)
		}
	}
}

// ConnectionsForIP is used by the handshake limiter (MAX_WS_PER_IP).
func (h *Hub) ConnectionsForIP(ip string) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.byIP[ip]
}

// Deliver pushes a frame to every live socket of an account. It reports
// whether at least one socket accepted the frame; otherwise the caller keeps
// the envelope in PostgreSQL for the offline pull path.
func (h *Hub) Deliver(accountID string, payload []byte) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	delivered := false
	for c := range h.byAccount[accountID] {
		select {
		case c.send <- payload:
			delivered = true
		default:
			h.log.Warn("slow websocket client, dropping push", "account", shortID(accountID))
		}
	}
	return delivered
}

func (h *Hub) CloseAll() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.closed = true
	for _, set := range h.byAccount {
		for c := range set {
			c.Close()
		}
	}
}

func (c *Client) Start(ctx context.Context) {
	go c.writePump(ctx)
}

func (c *Client) writePump(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case msg := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Send queues a frame without blocking. Returns false when the queue is full.
func (c *Client) Send(payload []byte) bool {
	select {
	case c.send <- payload:
		return true
	default:
		return false
	}
}

func (c *Client) Close() {
	c.closeOnce.Do(func() {
		close(c.done)
		c.hub.remove(c)
		_ = c.conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			time.Now().Add(writeWait))
		_ = c.conn.Close()
	})
}

// ------------------------------------------------------------ frame schema

// EnvelopeFrame is the only message type the server proactively pushes.
type EnvelopeFrame struct {
	Type        string    `json:"type"`
	ID          int64     `json:"id"`
	From        string    `json:"from,omitempty"`
	CipherType  int       `json:"cipher_type"`
	Body        string    `json:"body"`
	ClientMsgID string    `json:"client_msg_id"`
	SentAt      time.Time `json:"sent_at"`
}

type ReadyFrame struct {
	Type      string `json:"type"`
	AccountID string `json:"account_id"`
	ServerTS  int64  `json:"server_ts"`
}

func shortID(id string) string {
	if len(id) <= 10 {
		return id
	}
	return id[:10] + "…"
}
```

### `server/migrations/0001_init.sql`

```sql
-- Paperplane schema. No plaintext, no reversible username, no long-term ciphertext.
CREATE TABLE IF NOT EXISTS accounts (
    account_id   TEXT PRIMARY KEY,
    auth_hash    TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_seen_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS prekey_bundles (
    account_id             TEXT PRIMARY KEY,
    identity_key           BYTEA NOT NULL,
    signed_prekey_id       BIGINT NOT NULL,
    signed_prekey_public   BYTEA NOT NULL,
    signed_prekey_signature BYTEA NOT NULL,
    kyber_prekey_id        BIGINT,
    kyber_prekey_public    BYTEA,
    kyber_prekey_signature BYTEA,
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS one_time_prekeys (
    account_id TEXT NOT NULL,
    prekey_id  BIGINT NOT NULL,
    public_key BYTEA NOT NULL,
    PRIMARY KEY (account_id, prekey_id)
);

CREATE TABLE IF NOT EXISTS envelopes (
    id            BIGSERIAL PRIMARY KEY,
    target_id     TEXT NOT NULL,
    sender_id     TEXT,
    cipher_type   SMALLINT NOT NULL,
    body          BYTEA NOT NULL,
    client_msg_id TEXT NOT NULL,
    sent_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at    TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS envelopes_target_idx ON envelopes (target_id, id);

CREATE INDEX IF NOT EXISTS envelopes_expiry_idx ON envelopes (expires_at);

CREATE INDEX IF NOT EXISTS one_time_prekeys_account_idx ON one_time_prekeys (account_id);
```

### `xdp/Makefile`

```makefile
# Paperplane XDP filter build
# Requirements: clang (bpf target), libbpf-dev, bpftool, libelf-dev, root/CAP_BPF
CLANG   ?= clang
BPFTOOL ?= bpftool
CC      ?= g++
CFLAGS  ?= -g -O2 -Wall
BPF_CFLAGS ?= -target bpf -D__TARGET_ARCH_x86 -O2 -g -Wall -Werror

all: xdp_filter.bpf.o xdp_loader

xdp_filter.bpf.o: src/xdp_filter.c
	$(CLANG) $(BPF_CFLAGS) -c $< -o $@

xdp_loader: src/xdp_loader.cpp xdp_filter.bpf.o
	$(CC) $(CFLAGS) -std=c++17 $< -lbpf -lelf -o $@

vmlinux:
	$(BPFTOOL) btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h

check:
	bpftool feature probe | grep -q "eXtended DataPath" || echo "XDP not available on this host"

clean:
	rm -f xdp_filter.bpf.o xdp_loader vmlinux.h

.PHONY: all clean check
```

### `xdp/config.env`

```dotenv
XDP_ENABLED=true
XDP_IFACE=eth0
BAN_FILE=/var/lib/paperplane/bans.txt
BAN_POLL_SECONDS=5
```

### `xdp/src/xdp_filter.c`

```c
// SPDX-License-Identifier: AGPL-3.0
// Paperplane XDP packet filter: drop banned source IPs before the socket layer.
// Build: clang -target bpf -O2 -g -c xdp_filter.c -o xdp_filter.bpf.o
// Requires CONFIG_BPF_LSM, XDP socket support; run on a real NIC (not veth).
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <bpf/bpf_helpers.h>

struct lpm_key {
    __u32 prefixlen; /* network byte order length for LPM trie */
    __u32 addr;
};

/* LPM_TRIE: longest-prefix match ban table, populated by the loader. */
struct {
    __uint(type, BPF_MAP_TYPE_LPM_TRIE);
    __uint(max_entries, 1024);
    __uint(key_size, sizeof(struct lpm_key));
    __uint(value_size, sizeof(__u64));
    __uint(map_flags, BPF_F_NO_PREALLOC);
} ban_map SEC(".maps");

/* Per-CPU drop counter for observability. */
struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __uint(max_entries, 1);
    __uint(key_size, sizeof(__u32));
    __uint(value_size, sizeof(__u64));
} stats_map SEC(".maps");

static __always_inline int is_banned(__u32 addr) {
    struct lpm_key key = {
        .prefixlen = 32,
        .addr = addr, /* already in network byte order */
    };
    return bpf_map_lookup_elem(&ban_map, &key) != NULL;
}

SEC("xdp")
int xdp_paperplane_filter(struct xdp_md *ctx) {
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;

    struct ethhdr *eth = data;
    if ((void *)(eth + 1) > data_end)
        return XDP_PASS;

    /* UDP/TCP/ICMP v4 only; everything else bypasses. */
    if (eth->h_proto != __constant_htons(ETH_P_IP))
        return XDP_PASS;

    struct iphdr *ip = (void *)(eth + 1);
    if ((void *)(ip + 1) > data_end)
        return XDP_PASS;

    if (is_banned(ip->saddr)) {
        __u32 idx = 0;
        __u64 *cnt = bpf_map_lookup_elem(&stats_map, &idx);
        if (cnt)
            __sync_fetch_and_add(cnt, 1);
        return XDP_DROP;
    }
    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
```

### `xdp/src/xdp_loader.cpp`

```cpp
// SPDX-License-Identifier: AGPL-3.0
// Paperplane XDP loader: RAII ownership over libbpf objects, explicit detach on
// SIGTERM/SIGINT, and a periodic diff of /var/lib/paperplane/bans.txt into the
// LPM_TRIE map. Build: g++ -O2 -std=c++17 xdp_loader.cpp -lbpf -lelf -o xdp_loader
#include <atomic>
#include <csignal>
#include <cstdint>
#include <cstdio>
#include <cstring>
#include <fstream>
#include <iostream>
#include <net/if.h>
#include <string>
#include <thread>
#include <unordered_map>
#include <vector>

#include <bpf/bpf.h>
#include <bpf/libbpf.h>
#include <arpa/inet.h>

namespace {
std::atomic<bool> g_stop{false};
int g_bpf_obj_fd = -1;
int g_prog_fd = -1;

void on_signal(int) {
    g_stop = true;
}

// RAII guard for bpf_object: closes handle on scope exit.
class BpfObject {
  public:
    explicit BpfObject(bpf_object *o) : obj_(o) {}
    ~BpfObject() { if (obj_) bpf_object__close(obj_); }
    BpfObject(const BpfObject &) = delete;
    BpfObject &operator=(const BpfObject &) = delete;
    bpf_object *get() { return obj_; }

  private:
    bpf_object *obj_ = nullptr;
};

bool parseBanFile(const std::string &path, std::unordered_map<uint32_t, time_t> *out) {
    std::ifstream in(path);
    if (!in) return false;
    out->clear();
    std::string line;
    while (std::getline(in, line)) {
        if (line.empty() || line[0] == '#') continue;
        auto sp = line.find(' ');
        std::string ip = line.substr(0, sp);
        time_t until = 0;
        if (sp != std::string::npos) until = std::stoll(line.substr(sp + 1));
        in_addr addr{};
        if (inet_pton(AF_INET, ip.c_str(), &addr) != 1) continue;
        (*out)[addr.s_addr] = until;
    }
    return true;
}
}  // namespace

int main(int argc, char **argv) {
    if (argc < 3) {
        std::fprintf(stderr,
                     "usage: %s <iface> <bans.txt> [poll_seconds]\n"
                     "  -d   drop mode (default skb/generic fallback)\n",
                     argv[0]);
        return 1;
    }
    std::string iface = argv[1];
    std::string banFile = argv[2];
    int poll = argc > 3 ? std::atoi(argv[3]) : 5;

    std::signal(SIGINT, on_signal);
    std::signal(SIGTERM, on_signal);
    std::signal(SIGHUP, on_signal);

    int ifindex = if_nametoindex(iface.c_str());
    if (ifindex == 0) {
        std::fprintf(stderr, "interface %s not found\n", iface.c_str());
        return 1;
    }

    // Load BPF object.
    BpfObject obj(bpf_object__open("xdp_filter.bpf.o"));
    if (!obj.get()) {
        std::fprintf(stderr, "bpf_object__open failed\n");
        return 1;
    }
    if (bpf_object__load(obj.get())) {
        std::fprintf(stderr, "bpf_object__load failed (missing libbpf headers?)\n");
        return 1;
    }
    bpf_program *prog = bpf_object__find_program_by_name(obj.get(), "xdp_paperplane_filter");
    if (!prog) {
        std::fprintf(stderr, "program not found in object\n");
        return 1;
    }
    g_prog_fd = bpf_program__fd(prog);

    // Attach: try generic/skb first (safe on virtio/QEMU), then native.
    int flags = XDP_FLAGS_SKB_MODE;
    if (bpf_xdp_attach(ifindex, g_prog_fd, flags, nullptr) < 0) {
        std::fprintf(stderr, "skb attach failed, trying drv mode\n");
        flags = XDP_FLAGS_DRV_MODE;
        if (bpf_xdp_attach(ifindex, g_prog_fd, flags, nullptr) < 0) {
            std::perror("bpf_xdp_attach");
            return 1;
        }
    }
    std::printf("xdp attached on %s (flags=%d)\n", iface.c_str(), flags);

    int banMap = bpf_object__find_map_fd_by_name(obj.get(), "ban_map");

    // Periodic sync: diff file content against the BPF map.
    std::unordered_map<uint32_t, time_t> current;
    while (!g_stop) {
        std::unordered_map<uint32_t, time_t> wanted;
        if (parseBanFile(banFile, &wanted)) {
            for (const auto &[ip, until] : wanted) {
                if (current.find(ip) == current.end()) {
                    if (bpf_map_update_elem(banMap, &(struct { uint32_t p = 32; uint32_t a; }){32, ip},
                                            &(uint64_t{1}), BPF_ANY)) {
                        /* ignore transient errors */
                    }
                }
            }
            for (const auto &[ip, _] : current) {
                if (wanted.find(ip) == wanted.end()) {
                    struct { uint32_t p = 32; uint32_t a; } key{32, ip};
                    bpf_map_delete_elem(banMap, &key);
                }
            }
            current = wanted;
        }
        std::this_thread::sleep_for(std::chrono::seconds(poll));
    }

    // Explicit detach on shutdown.
    if (bpf_xdp_detach(ifindex, flags, nullptr) < 0) {
        std::perror("bpf_xdp_detach");
    } else {
        std::printf("xdp detached cleanly\n");
    }
    return 0;
}
```
