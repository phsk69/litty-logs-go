# litty-logs-go — the most bussin Go logging library no cap 🔥

# build the whole thing bestie 🏗️🔥
build:
    go build ./...

# run all tests — check them vibes 🧪🔥
test *args:
    go test -v -race -count=1 ./... {{args}}

# lint the code — vet for that clean energy 🔍
vet:
    go vet ./...

# run the basic example — see litty-logs in action bestie 🔥
example:
    go run ./examples/basic/

# bump the version bestie — usage: just bump major|minor|patch 🔥
bump part:
    #!/usr/bin/env bash
    set -euo pipefail
    current=$(cat VERSION)
    IFS='.' read -r major minor patch <<< "${current%%-*}"
    case "{{part}}" in
        major) major=$((major + 1)); minor=0; patch=0 ;;
        minor) minor=$((minor + 1)); patch=0 ;;
        patch) patch=$((patch + 1)) ;;
        *) echo "fam thats not a valid bump part — use major, minor, or patch no cap 😤"; exit 1 ;;
    esac
    new_version="${major}.${minor}.${patch}"
    echo -n "${new_version}" > VERSION
    echo "version went from ${current} -> ${new_version} lets gooo 🔥"

# add a pre-release label — usage: just bump-pre dev.1 🏷️
bump-pre label:
    #!/usr/bin/env bash
    set -euo pipefail
    current=$(cat VERSION)
    base="${current%%-*}"
    new_version="${base}-{{label}}"
    echo -n "${new_version}" > VERSION
    echo "version went from ${current} -> ${new_version} (pre-release vibes) 🏷️"

# gitflow release — start branch, bump, finish, push 🚀
release part:
    #!/usr/bin/env bash
    set -euo pipefail
    if [ -n "$(git status --porcelain)" ]; then
        echo "fam your working tree is dirty, commit or stash first no cap 😤"
        exit 1
    fi
    current=$(cat VERSION)
    IFS='.' read -r major minor patch <<< "${current%%-*}"
    case "{{part}}" in
        major) major=$((major + 1)); minor=0; patch=0 ;;
        minor) minor=$((minor + 1)); patch=0 ;;
        patch) patch=$((patch + 1)) ;;
        *) echo "fam thats not a valid bump part — use major, minor, or patch no cap 😤"; exit 1 ;;
    esac
    new_version="${major}.${minor}.${patch}"
    echo "starting the gitflow release ritual bestie 🕯️"
    echo "  ${current} -> ${new_version}"
    echo ""
    git flow release start "v${new_version}"
    echo -n "${new_version}" > VERSION
    git add VERSION
    git commit -m "bump: v${new_version} incoming no cap 🔥"
    GIT_MERGE_AUTOEDIT=no git flow release finish "v${new_version}" -m "v${new_version} dropped no cap 🔥"
    echo ""
    echo "release v${new_version} complete 🔥"
    echo "pushing develop, main, and tag to origin 📤"
    git push origin develop main "v${new_version}"

# release the current version as-is without bumping 🚀
release-current:
    #!/usr/bin/env bash
    set -euo pipefail
    if [ -n "$(git status --porcelain)" ]; then
        echo "fam your working tree is dirty, commit or stash first no cap 😤"
        exit 1
    fi
    version=$(cat VERSION)
    echo "releasing v${version} as-is bestie 🕯️"
    git flow release start "v${version}"
    GIT_MERGE_AUTOEDIT=no git flow release finish "v${version}" -m "v${version} dropped no cap 🔥"
    echo "release v${version} complete 🔥"
    git push origin develop main "v${version}"

# finish whatever gitflow branch youre on 🏁
finish:
    #!/usr/bin/env bash
    set -euo pipefail
    branch=$(git rev-parse --abbrev-ref HEAD)
    if [ -n "$(git status --porcelain)" ]; then
        echo "fam your working tree is dirty, commit or stash first no cap 😤"
        exit 1
    fi
    if [[ "$branch" == hotfix/* ]]; then
        version="${branch#hotfix/}"; kind="hotfix"
    elif [[ "$branch" == release/* ]]; then
        version="${branch#release/}"; kind="release"
    else
        echo "bruh youre on '${branch}' — thats not a hotfix or release branch 💀"
        exit 1
    fi
    version_clean="${version#v}"
    GIT_MERGE_AUTOEDIT=no git flow "${kind}" finish "${version}" -m "v${version_clean} ${kind} dropped no cap 🔥"
    git push origin develop main "${version}"
    echo "v${version_clean} complete 🔥"
