import neopixel
import machine

# Wheel method from Adafruit Neopixel library examples (MIT license)
# https://github.com/adafruit/Adafruit_CircuitPython_NeoPixel/blob/master/examples/neopixel_simpletest.py
def wheel(pos):
    # Input a value 0 to 255 to get a color value.
    # The colours are a transition r - g - b - back to r.
    if pos < 0 or pos > 255:
        return (0, 0, 0, 0)
    if pos < 85:
        return (255 - pos * 3, pos * 3, 0, 0)
    if pos < 170:
        pos -= 85
        return (0, 255 - pos * 3, pos * 3, 0)
    pos -= 170
    return (pos * 3, 0, 255 - pos * 3, 0)

# There are 32 pixels in the strip attached to my monitor.
pixels = 32
px = neopixel.NeoPixel(machine.Pin(12), pixels, bpp=4) # sk6812 are RGBW, bpp=4 mode.

# Rainbow with every 4th pixel to be white.
# I find this helps make the lighting a bit brighter (helps for reading/soldering),
# but still has the rainbow aesthetic ;)
for i in range(pixels):
    if i % 4 == 0:
        px[i] = (0,0,0,160)
    else:
        r = wheel((i * 256 // pixels) & 255)
        px[i] = r

px.write()
