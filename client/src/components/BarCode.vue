<template>
  <div>
    <Dialog :dialog="dialog" target="alert">
      <template v-slot:headline>
        <h2 style="color:red;">エラー</h2>
      </template>
      <template v-slot:content>
        <h3>{{ errorMessage }}</h3>
      </template>
    </Dialog>
    <div>
      <video id="video" width="100%" height="100%" style="border: 1px solid gray"></video>
    </div>
    <div id="sourceSelectPanel" style="display:block" v-show="isSelect">
      <label for="sourceSelect">Change video source:</label>
      <select id="sourceSelect" style="max-width:400px"></select>
    </div>
  </div>
</template>

<script>
import { BrowserBarcodeReader } from '@zxing/library'
import Dialog from './shared/Dialog'
export default {
  name: 'BarCode',
  components: {
    Dialog
  },
  data () {
    return {
      selectedDeviceId: null,
      dialog: {
        isOpen: false,
        closeText: 'close',
        target: ''
      },
      errorMessage: 'エラー',
      isSelect: false
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
      this.codeReader
        .decodeFromVideoDevice(this.selectedDeviceId, 'video', result => {
          if (result !== null && result !== undefined) {
            if (this.checkDigit(result.text) && this.checkISBN(result.text)) {
              this.$emit('changeCode', result.text)
              this.$emit('search')
            }
          }
        })
        .then(result => {
          if (result !== null && result !== undefined) {
            document.getElementById('result').textContent = result.text
          }
        })
        .catch(err => {
          this.setAlert(err)
          document.getElementById('result').textContent = 'err'
        })
    },
    stop () {
      this.codeReader.reset()
    },
    setDialog (closeText, target) {
      this.dialog = {
        isOpen: true,
        closeText: closeText,
        target: target
      }
    },
    setAlert (errmsg) {
      this.errorMessage = errmsg
      this.setDialog('close', 'alert')
    }
  },
  mounted () {
    this.codeReader = new BrowserBarcodeReader()
    this.codeReader
      .getVideoInputDevices()
      .then(videoInputDevices => {
        const sourceSelect = document.getElementById('sourceSelect')
        this.selectedDeviceId = videoInputDevices[0].deviceId
        if (videoInputDevices.length > 1) {
          videoInputDevices.forEach(element => {
            const sourceOption = new Option(element.label, element.deviceId, true, true)
            sourceSelect.appendChild(sourceOption)
          })
          sourceSelect.onchange = () => {
            this.selectedDeviceId = sourceSelect.value
            this.stop()
            this.start()
          }
          this.isSelect = true
        }
      })
      .catch(err => {
        this.setAlert(err)
      })
    this.start()
  },
  destroyed () {
    this.codeReader.reset()
  }
}
</script>
