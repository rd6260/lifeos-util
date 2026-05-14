# lifeos-util

small programs and binaries that make the operating system feel like mine.

nothing here is necessary. that's the point.

---

## what this is

a collection of personal utilities — command-line tools, scripts, and binaries built for daily use. they don't solve big problems. they solve the problem of a system that feels generic, impersonal, too loud or not quite right.

some print a line of text. some might do more eventually. they're written to be run, forgotten about, and quietly appreciated when they surface again.

---

## tools

### `lifeos-util`

quiet little messages for quiet little moments.

```
lifeos-util [--flag]
```

| flag | description |
|---|---|
| `--time-whisper` | a soft, time-aware mood message — a few quiet words based on the hour |
| `--system-ambient` | a gentle message as if the system itself is speaking |

**examples**

```
$ lifeos-util --time-whisper
the good part of the day

$ lifeos-util --system-ambient
the kernel is unbothered — it usually is
```

---

## building

requires Go. you can build and install using the provided makefile.

```sh
make build
```

to install it to `~/.local/bin/`:

```sh
make install
```

---

## philosophy

these tools are personal. they're not meant to be useful to everyone — they're meant to be useful to one person's particular way of living inside a terminal. fork it, ignore it, or build on it.

the repo will grow slowly, as the need arises.

---

## license

do what you want with it.
