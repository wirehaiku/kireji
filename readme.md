# Kireji

**Kireji** is a plaintext note-taking CLI, written in [Go 1.20][gver] by [Stephen Malone][smal].

- See [`changes.md`][chng] for the complete changelog.
- See [`license.md`][lcns] for licensing information (BSD-3).

If you've got a directory full of plaintext notes you need to handle, Kireji can give you an easy-breezy command-line interface to do just that.
Just tell Kireji where your notes are and let it do the rest.

## Installation

You can install Kireji with `go get`...

```
go install github.com/wirehaiku/kireji@latest
```

or download a binary from the [latest release][rels].

## Configuration

Kireji just needs two environment variables to be set and exported.

```fish
# KIREJI_DIR is the path to your notes directory.
KIREJI_DIR = "~/notes"

# KIREJI_EXT is the file extension your notes use.
KIREJI_EXT = ".txt"
```

That's it! That's all the configuration you need.

## Commmands

### Syntax

Kireji represents note files as lowercase base names.
For example, the file `~/notes/foo.txt` is represented as `foo` and asking Kireji to make the note `FOO BAR` results in `~/notes/foo-bar.txt` being created.

### `edit NOTE`

Open an existing note in the default editor, according to the environment variable `EDITOR`.

<details> <summary>Example:</summary>

```
$ kireji edit party-planner
```

</details>

### `help [COMMAND]`

Print a help message for all commands or a specified command.

<details> <summary>Example:</summary>

```
$ kireji help edit
edit:
  Edit a new or existing note.
  $ kireji edit NOTE
```

</details>

### `junk NOTE`

"Delete" an existing note by changing its file extension to `.junk`.

<details> <summary>Example:</summary>

```
$ kireji junk old-notes
$ ls ~/notes
old-notes.junk
```

</details>

### `list [SUBSTRING]`

List all existing notes in your directory, or notes with names containing a substring.

<details> <summary>Example:</summary>

```
$ kireji list
2023-goals
groceries
party-planner

$ kireji list 2023
2023-goals
```

</details>

### `make NOTE`

Create a new empty note in your directory.

<details> <summary>Example:</summary>

```
$ kireji make new-note
$ ls ~/notes
new-note.txt
```

</details>

### `show NOTE`

Print an existing note's contents.

<details> <summary>Example:</summary>

```
$ kireji show groceries
Need to buy: bread, milk, tomato soup, party streamers...
```

</details> 

## Contributing

Please submit all feature requests and bug reports to the [issue tracker][bugs], thank you.

[bugs]: https://github.com/wirehaiku/kireji/issues
[chng]: https://github.com/wirehaiku/kireji/blob/main/changes.md
[gver]: https://go.dev/doc/go1.20
[lcns]: https://github.com/wirehaiku/kireji/blob/main/license.md
[rels]: https://github.com/wirehaiku/kireji/releases/latest
[smal]: https://wirehaiku.org/


