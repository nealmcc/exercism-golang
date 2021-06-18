package piglatin

import (
	"bufio"
	"log"
	"piglatin/pkg/trie"
	"strings"
)

func Sentence(text string) string {
	var b strings.Builder

	s := bufio.NewScanner(strings.NewReader(text))
	s.Split(bufio.ScanWords)

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}

		if b.Len() > 0 {
			b.WriteByte(' ')
		}

		word := s.Text()
		rule, _ := t.Evaluate(word)
		bytes := s.Bytes()
		b.Write(bytes[rule.N:])
		b.Write(bytes[:rule.N])
		b.WriteString("ay")
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return b.String()
}

var t trie.Trie = trie.WithRules(rules...)

var rules []trie.Rule = []trie.Rule{
	// a word begins with a vowel sound
	{Prefix: "a", N: 0},
	{Prefix: "e", N: 0},
	{Prefix: "i", N: 0},
	{Prefix: "o", N: 0},
	{Prefix: "u", N: 0},
	{Prefix: "xr", N: 0}, // xray
	{Prefix: "xm", N: 0}, // xmas
	{Prefix: "yt", N: 0}, // yttria

	// single consononant sound: (including silent letters)
	{Prefix: "p", N: 1},  // pin
	{Prefix: "b", N: 1},  // bin
	{Prefix: "t", N: 1},  // tin
	{Prefix: "pt", N: 2}, // pterodactyl
	{Prefix: "d", N: 1},  // dog
	{Prefix: "dh", N: 2}, // dharma
	{Prefix: "pd", N: 2}, // ?
	{Prefix: "k", N: 1},  // key
	{Prefix: "kh", N: 2}, // khan
	{Prefix: "q", N: 1},  // ?
	{Prefix: "c", N: 1},  // cat
	{Prefix: "g", N: 1},  // gate, gin
	{Prefix: "gh", N: 2}, // ghost
	{Prefix: "gn", N: 2}, // gnome
	{Prefix: "m", N: 1},  // mate
	{Prefix: "n", N: 1},  // no
	{Prefix: "kn", N: 2}, // knuckle
	{Prefix: "mn", N: 2}, // mnemonic
	{Prefix: "pn", N: 2}, // pneumatic
	{Prefix: "f", N: 1},  // fin
	{Prefix: "ph", N: 2}, // phenomenal
	{Prefix: "v", N: 1},  // van
	{Prefix: "th", N: 2}, // thin, the
	{Prefix: "s", N: 1},  // sun
	{Prefix: "ps", N: 2}, // psyche
	{Prefix: "z", N: 1},  // zoo
	{Prefix: "x", N: 1},  // xylophone
	{Prefix: "sh", N: 2}, // shop
	{Prefix: "ch", N: 2}, // chop
	{Prefix: "j", N: 1},  // job
	{Prefix: "w", N: 1},  // win
	{Prefix: "r", N: 1},  // run
	{Prefix: "rh", N: 2}, // rhyme
	{Prefix: "l", N: 1},  // lip
	{Prefix: "ll", N: 2}, // llama
	{Prefix: "y", N: 1},  // yellow
	{Prefix: "h", N: 1},  // how
	{Prefix: "wh", N: 2}, // who

	// two-consonant blends
	{Prefix: "sm", N: 2},  // small
	{Prefix: "sn", N: 2},  // snap
	{Prefix: "st", N: 2},  // stall
	{Prefix: "sw", N: 2},  // swell
	{Prefix: "sv", N: 2},  // svelte
	{Prefix: "zw", N: 2},  // zwieback
	{Prefix: "sk", N: 2},  // skald
	{Prefix: "sc", N: 2},  // scuttle
	{Prefix: "sch", N: 3}, // school
	{Prefix: "sl", N: 2},  // slow
	{Prefix: "sr", N: 2},  // ?
	{Prefix: "vr", N: 2},  // vroom
	{Prefix: "vl", N: 2},  // vlad ?
	{Prefix: "sp", N: 2},  // spot
	{Prefix: "sph", N: 3}, // sphere
	{Prefix: "sf", N: 2},  // ?
	{Prefix: "tw", N: 2},  // twig
	{Prefix: "thr", N: 3}, // throw
	{Prefix: "thw", N: 3}, // thwart
	{Prefix: "dr", N: 2},  // drip
	{Prefix: "tr", N: 2},  // trip
	{Prefix: "dw", N: 2},  // dwell
	{Prefix: "tw", N: 2},  // twin
	{Prefix: "qu", N: 2},  // queen
	{Prefix: "kw", N: 2},  // ?
	{Prefix: "pl", N: 2},  // play
	{Prefix: "pr", N: 2},  // pray
	{Prefix: "bl", N: 2},  // blue
	{Prefix: "br", N: 2},  // bray
	{Prefix: "gl", N: 2},  // glow
	{Prefix: "gr", N: 2},  // green
	{Prefix: "fl", N: 2},  // fly
	{Prefix: "fr", N: 2},  // free
	{Prefix: "phl", N: 3}, // phlegm
	{Prefix: "phr", N: 3}, // phrase
	{Prefix: "cr", N: 2},  // crunch
	{Prefix: "cl", N: 2},  // clay
	{Prefix: "kl", N: 2},  // ? Klingon
	{Prefix: "kr", N: 2},  // ?
	{Prefix: "chr", N: 3}, // chronic
	{Prefix: "chl", N: 3}, // chlorophyll
	{Prefix: "wr", N: 2},  // wrong
	{Prefix: "wl", N: 2},  // ?
	{Prefix: "shr", N: 3}, // shred
	{Prefix: "shm", N: 3}, // shmooze ?
	{Prefix: "ts", N: 2},  // tsar
	{Prefix: "tz", N: 2},  // tzar
	{Prefix: "cz", N: 2},  // czar
	{Prefix: "fj", N: 2},  // fjord
	{Prefix: "dj", N: 2},  // djinn

	// three-consonant blends:
	{Prefix: "spl", N: 3},  // split
	{Prefix: "spr", N: 3},  // sprint
	{Prefix: "str", N: 3},  // string
	{Prefix: "skl", N: 3},  // ?
	{Prefix: "scl", N: 3},  // sclerosis
	{Prefix: "skr", N: 3},  // ?
	{Prefix: "scr", N: 3},  // screen
	{Prefix: "skw", N: 3},  // ?
	{Prefix: "squ", N: 3},  // squid
	{Prefix: "schn", N: 4}, // schnapps
	{Prefix: "schl", N: 4}, // schlep
	{Prefix: "schm", N: 4}, // schmooze ?
}
