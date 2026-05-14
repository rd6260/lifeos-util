package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Flags — more can be added below as the project grows
	whisper := flag.Bool("time-whisper", false, "a soft, time-aware mood message — a few quiet words based on the hour")
	ambient := flag.Bool("system-ambient", false, "a gentle message as if the system itself is speaking")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "lifeos-util — quiet little messages for quiet little moments\n\n")
		fmt.Fprintf(os.Stderr, "usage:\n")
		fmt.Fprintf(os.Stderr, "  lifeos-util [--flag]\n\n")
		fmt.Fprintf(os.Stderr, "flags:\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  --%s\n\t%s\n", f.Name, f.Usage)
		})
	}

	flag.Parse()

	// Seed randomness
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch {
	case *whisper:
		fmt.Println(pickWhisper(rng))
	case *ambient:
		fmt.Println(pickAmbient(rng))
	default:
		flag.Usage()
		os.Exit(0)
	}
}

// pickWhisper returns a cozy, time-aware mood message.
func pickWhisper(rng *rand.Rand) string {
	hour := time.Now().Hour()

	var pool []string

	switch {
	case hour >= 0 && hour < 4:
		// Deep night
		pool = []string{
			"somewhere past midnight",
			"still awake?",
			"just you now",
			"quiet hours",
			"the night goes long",
			"this late, things feel different",
			"very late or very early",
			"the dark, unbroken",
			"not sleeping yet",
			"the small hours",
			"hours no one names",
		}
	case hour >= 4 && hour < 7:
		// Pre-dawn
		pool = []string{
			"the dark is softening",
			"almost morning",
			"still deciding",
			"between night and day",
			"the first birds soon",
			"early light, barely",
			"the world not yet up",
			"pre-dawn quiet",
			"cold and still out there",
			"the sky, slowly",
		}
	case hour >= 7 && hour < 12:
		// Morning
		pool = []string{
			"take your time",
			"still young, the morning",
			"slow start — that's fine",
			"coffee weather",
			"no rush yet",
			"the day, just beginning",
			"easy does it",
			"morning light, clean",
			"plenty of time still",
			"not quite afternoon",
			"the good part of the day",
		}
	case hour >= 12 && hour < 15:
		// Midday
		pool = []string{
			"the afternoon settles in",
			"peak hours — breathe",
			"midday quiet, somehow",
			"sun at its highest",
			"half the day gone",
			"the long middle",
			"full daylight",
			"everything lit up",
			"lunch, maybe",
			"the day at its loudest",
		}
	case hour >= 15 && hour < 18:
		// Late afternoon
		pool = []string{
			"the light's changing",
			"slow afternoon",
			"golden-hour feeling",
			"winding down, eventually",
			"late afternoon, the gentle kind",
			"the day tilting",
			"shadows getting longer",
			"almost evening",
			"the soft slant of it",
			"quieter now than noon",
		}
	case hour >= 18 && hour < 21:
		// Evening
		pool = []string{
			"slow evening",
			"the day folding up",
			"soft and low, the light",
			"somewhere, dinner",
			"the world slows here",
			"evening settling in",
			"the good kind of tired",
			"clocking out, slowly",
			"the day almost done",
			"dimmer now",
			"the end of it, nearly",
		}
	default:
		// Night (21–23)
		pool = []string{
			"quiet hours",
			"night, full swing",
			"the late crowd",
			"almost midnight",
			"still going?",
			"the night, deep now",
			"everyone else is asleep",
			"running late, or just running",
			"the long end of the day",
			"near the close of it",
		}
	}

	return pool[rng.Intn(len(pool))]
}

// pickAmbient returns a message as if the system itself is speaking.
func pickAmbient(rng *rand.Rand) string {
	pool := []string{
		"the rain outside sounds softer tonight",
		"the air feels lighter here",
		"everything hums quietly tonight",
		"the world sounds distant for a while",
		"the static feels comforting somehow",
		"something in the room shifted, gently",
		"the hum of things running in the background",
		"the fans spin slower when no one's watching",
		"signals drift in from far away",
		"the lights are a little dimmer on purpose",
		"somewhere, a server blinks in the dark",
		"bits and silence, mostly silence",
		"the wind has no opinion — neither does the machine",
		"the cursor blinks at its own pace",
		"the room holds still for a moment",
		"nothing urgent, nothing needs to be",
		"all the open tabs, resting",
		"the network is quiet — the thoughts aren't",
		"things are running, just barely",
		"it's warm in here from the processor",
		"a low hum, a familiar one",
		"the hard drive sighs, just slightly",
		"somewhere downstream, something waits",
		"the uptime keeps climbing, silently",
		"packets moving, no one watching",
		"the machine has been patient all day",
		"a process completes somewhere you'll never see",
		"the logs scroll by — nothing unusual",
		"somewhere a cron job woke and went back to sleep",
		"the buffer empties without ceremony",
		"the clock ticks at the hardware level, steady",
		"memory frees itself when you're not looking",
		"something is always compiling somewhere",
		"the system breathes in its own way",
		"a thread waits, politely",
		"the kernel is unbothered — it usually is",
		"idle cycles, spent on nothing, contentedly",
		"the connection holds, for now",
		"no alerts — just the ordinary hum of things",
		"the terminal remembers every command you forgot",
	}

	return pool[rng.Intn(len(pool))]
}
