<template>
  <div>
    <div>
      <video id="video" width="100%" height="100%" style="border: 1px solid gray"></video>
    </div>
    <div id="sourceSelectPanel" style="display:none">
      <label for="sourceSelect">Change video source:</label>
      <select id="sourceSelect" style="max-width:400px"></select>
    </div>
  </div>
</template>

<script>
import { BrowserBarcodeReader } from '@zxing/library'
export default {
  name: 'BarCode',
  data () {
    return {
      selectedDeviceId: null
    }
  },
  methods: {
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
    },
    start () {
      console.log(this.codeReader)
      this.codeReader
        .decodeFromVideoDevice(this.selectedDeviceId, 'video', result => {
          console.log(result)
          if (result !== null) {
            if (this.checkDigit(result.text) && this.checkISBN(result.text)) {
              this.$emit('changeCode', result.text)
              this.$emit('search')
            }
          }
        })
        .then(result => {
          console.log(result)
          document.getElementById('result').textContent = result.text
        })
        .catch(err => {
          console.error(err, 'error')
          document.getElementById('result').textContent = 'err'
        })
      console.log(
        `Started continous decode from camera with id ${this.selectedDeviceId}`
      )
    },
    stop () {
      this.codeReader.reset()
    }
  },
  computed: {
    computedWidth () {
      let width = 100
      console.log(window.innerWidth * 0.75)
      return {
        width: `${width}px`,
        height: `${width * 0.75}px`
      }
    }
  },
  mounted () {
    this.codeReader = new BrowserBarcodeReader()
    console.log('ZXing code reader initialized')
    this.codeReader
      .getVideoInputDevices()
      .then(videoInputDevices => {
        const sourceSelect = document.getElementById('sourceSelect')
        this.selectedDeviceId = videoInputDevices[0].deviceId
        if (videoInputDevices.length > 1) {
          videoInputDevices.forEach(element => {
            const sourceOption = document.createElement('option')
            sourceOption.text = element.label
            sourceOption.value = element.deviceId
            sourceSelect.appendChild(sourceOption)
          })
          sourceSelect.onchange = () => {
            this.selectedDeviceId = sourceSelect.value
          }
          const sourceSelectPanel = document.getElementById(
            'sourceSelectPanel'
          )
          sourceSelectPanel.style.display = 'block'
        }
      })
      .catch(err => {
        console.error(err)
      })
    this.start()
  },
  destroyed () {
    this.codeReader.reset()
  }
}
</script>
