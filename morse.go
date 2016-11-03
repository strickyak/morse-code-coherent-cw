/* github.com/strickyak/morse-code-coherent-cw

Input words as os.Args.
Output CCW morse code (dit=100ms, 800Hz)
with 8000 one-byte samples/second for /dev/audio.

  $ go run morse.go  vvv vvv cq cq de kph > /dev/audio

  $ go run morse.go  vvv vvv cq cq de kph | gst-launch-0.10 -v -m filesrc location=/dev/stdin ! audio/x-raw-int, endianness=4321, channels=1, rate=8000, width=8, depth=8, signed=true ! audioconvert ! audioresample ! alsasink

References:
http://www.sigidwiki.com/wiki/Coherent_CW
http://www.nu-ware.com/NuCode%20Help/index.html?morse_code_structure_and_timing_.htm

Hint to get your Linux /dev/audio back:
  modprobe snd_pcm_oss
  modprobe snd_seq_oss
  modprobe snd_mixer_oss
Thank You, Kilian Foth:
http://unix.stackexchange.com/questions/120701/what-is-the-modern-equivalent-of-reading-dev-audio

*/
package main

import (
  "os"
  "strings"
)

const speed = 1.0  // Use 1.0 for normal 12wpm CCW.

var Code = map[rune]string {
  'a': ".-",
  'b': "-...",
  'c': "-.-.",
  'd': "-..",
  'e': ".",
  'f': "..-.",
  'g': "--.",
  'h': "....",
  'i': "..",
  'j': ".---",
  'k': "-.-",
  'l': ".-..",
  'm': "--",
  'n': "-.",
  'o': "---",
  'p': ".--.",
  'q': "--.-",
  'r': ".-.",
  's': "...",
  't': "-",
  'u': "..-",
  'v': "...-",
  'w': ".--",
  'x': "-..-",
  'y': "-.--",
  'z': "--..",
  '1': ".----",
  '2': "..---",
  '3': "...--",
  '4': "....-",
  '5': ".....",
  '6': "-....",
  '7': "--...",
  '8': "---..",
  '9': "----.",
  '0': "-----",
  '/': "-..-.",
  '.': ".-.-.-",
  ',': "--..--",
}

const Len800Hz = 8000 / 800
var Wave800Hz = [Len800Hz]byte {
  0, 60, 100, 100, 60, 0, -60+256, -100+256, -100+256, -60+256,
}

func Millis(tone bool, ms int) {
  numSamples := int(float64(ms * 8) / speed)
  buf := make([]byte, numSamples)
  if tone {
    for i := 0; i < numSamples; i++ {
      buf[i] = Wave800Hz[i%Len800Hz]
    }
  }
  os.Stdout.Write(buf)
}

func Dit() {
  Millis(true, 100)
  Millis(false, 100)
}
func Dah() {
  Millis(true, 300)
  Millis(false, 100)
}

func Vocalize(s string) {
  for _, r := range strings.ToLower(s) {
    x, ok := Code[r]
    if !ok {
      x = Code['/']  // Use '/' for unknown chars.
    }
    for _, c := range x {
      switch c {
        case '.': Dit()
        case '-': Dah()
        default: panic(x)
      }
    }
    Millis(false, 200)
  }
  Millis(false, 400)
}

func main() {
  for _, w := range os.Args[1:] {
    Vocalize(w)
  }
}
