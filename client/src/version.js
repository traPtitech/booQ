const production = 'v1.0.2' // バージョンアップの時はここを書き換えてください

let version = ''
switch (process.env.VUE_APP_API_ENDPOINT) {
  case 'http://localhost:3000':
    version = 'dev'
    break
  case 'https://booq-dev.tokyotech.org':
    version = 'staging'
    break
  default:
    version = production
    break
}
export default version
