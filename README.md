# morse-code-coherent-cw

Generates alaw morse code audio at 12 WPM with 100ms/dit Coherent CW

* Input words as os.Args.
* Output CCW morse code (dit=100ms, 800Hz)
* with 8000 one-byte samples/second for /dev/audio.

```
  $ go run morse.go  vvv vvv cq cq de kph > /dev/audio

  $ go run morse.go  vvv vvv cq cq de kph | gst-launch-0.10 -v -m filesrc location=/dev/stdin ! audio/x-raw-int, endianness=4321, channels=1, rate=8000, width=8, depth=8, signed=true ! audioconvert ! audioresample ! alsasink
```

### References:
* http://www.sigidwiki.com/wiki/Coherent_CW
* http://www.nu-ware.com/NuCode%20Help/index.html?morse_code_structure_and_timing_.htm

### Hint to get your Linux /dev/audio back:

```
    modprobe snd_pcm_oss
    modprobe snd_seq_oss
    modprobe snd_mixer_oss
```

### Thank You, Kilian Foth:
* http://unix.stackexchange.com/questions/120701/what-is-the-modern-equivalent-of-reading-dev-audio
