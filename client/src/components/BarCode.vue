<template>
  <div id="interactive" class="viewport scanner" :style="styles">
    <video class="quagga"/>
    <canvas class="drawingBuffer quagga" />
  </div>
</template>

<script>
import Quagga from 'quagga'
export default {
  name: 'QuaggaScanner',
  props: {
    onProcessed: {
      type: Function,
      default (result) {
        let drawingCtx = Quagga.canvas.ctx.overlay
        let drawingCanvas = Quagga.canvas.dom.overlay
        if (result) {
          if (result.boxes) {
            drawingCtx.clearRect(
              0,
              0,
              parseInt(drawingCanvas.getAttribute('width')),
              parseInt(drawingCanvas.getAttribute('height'))
            )
            result.boxes
              .filter(function (box) {
                return box !== result.box
              })
              .forEach(function (box) {
                Quagga.ImageDebug.drawPath(box, { x: 0, y: 1 }, drawingCtx, {
                  color: 'green',
                  lineWidth: 2
                })
              })
          }
          if (result.box) {
            Quagga.ImageDebug.drawPath(result.box, { x: 0, y: 1 }, drawingCtx, {
              color: '#00F',
              lineWidth: 2
            })
          }
          if (result.codeResult && result.codeResult.code) {
            Quagga.ImageDebug.drawPath(
              result.line,
              { x: 'x', y: 'y' },
              drawingCtx,
              { color: 'red', lineWidth: 3 }
            )
          }
        }
      }
    },
    readerTypes: {
      type: Array,
      default: () => ['ean_reader']
    },
    readerSize: {
      type: Object,
      default: () => ({
        width: 320,
        height: 240
      }),
      validator: o =>
        typeof o.width === 'number' && typeof o.height === 'number'
    },
    aspectRatio: {
      type: Object,
      default: () => ({
        min: 1,
        max: 2
      }),
      validator: o => typeof o.min === 'number' && typeof o.max === 'number'
    }
  },
  data () {
    return {
      quaggaState: {
        inputStream: {
          type: 'LiveStream',
          constraints: {
            width: 640,
            height: 480,
            facingMode: 'environment',
            aspectRatio: { min: 1, max: 2 }
          }
        },
        locator: {
          patchSize: 'medium',
          halfSample: true
        },
        numOfWorkers: 2,
        frequency: 10,
        decoder: {
          readers: this.readerTypes
        },
        locate: true
      }
      // styles: {
      //   display: 'inlineBlock',
      //   width: `${this.quaggaState.inputStream.constraints.width}px`,
      //   height: `${this.quaggaState.inputStream.constraints.width * 0.75}px`
      // }
    }
  },
  mounted () {
    const displaySize = document.getElementById('barcodewrapper').clientWidth
    const width = Math.min(displaySize, 640)
    this.quaggaState.inputStream.constraints.width = width
    this.quaggaState.inputStream.constraints.height = width * 0.75
    Quagga.init(this.quaggaState, function (err) {
      if (err) {
        return console.error(err)
      }
      Quagga.start()
    })
    Quagga.onDetected(this.onDetected)
    Quagga.onProcessed(this.onProcessed)
  },
  computed: {
    styles () {
      return {
        width: `${this.quaggaState.inputStream.constraints.width}px`,
        height: `${this.quaggaState.inputStream.constraints.width * 0.75}px`
      }
    }
  },
  methods: {
    onDetected (data) {
      if (this.checkDigit(data.codeResult.code) && this.checkISBN(data.codeResult.code)) {
        this.$emit('changeCode', data.codeResult.code)
        this.$emit('search')
      }
    },
    checkDigit (isbn) {
      const arrIsbn = isbn
        .toString()
        .split('')
        .map(num => parseInt(num))
      let remainder = 0
      const Digit = arrIsbn.pop()

      arrIsbn.forEach((num, index) => {
        remainder += num * (index % 2 === 0 ? 1 : 3)
      })
      remainder %= 10
      remainder = remainder === 0 ? 0 : 10 - remainder

      return Digit === remainder
    },
    checkISBN (isbn) {
      return isbn.slice(0, 3) === '978' || isbn.slice(0, 3) === '979'
    }
  },
  destroyed () {
    Quagga.stop()
  }
}
</script>

<style scoped>
.viewport {
  position: relative;
}
.viewport canvas,
.viewport video {
  position: absolute;
  left: 0;
  top: 0;
}
.quagga-wrapper {
  display: inline-block;
  width: 640px;
  height: 480px;
}
.quagga {
  display: inline-block;
  left: 0;
  right:0;
  margin: auto;
}
</style>
