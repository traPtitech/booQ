<template>
  <ItemList :items="data" />
  <!-- <div>{{data}}</div> -->
</template>

<script>
import ItemList from './shared/ItemList'
import axios from 'axios'

export default {
  name: 'AllItemPage',
  components: {
    ItemList
  },
  data () {
    return {
      data: null,
      error: null,
      sampleData: [
        {
          'id': 1,
          'name': 'プロジェクター',
          'code': 9784041026403,
          'type': 0,
          'owners': [
            {
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': true
            }
          ],
          'latest_logs': [
            {
              'id': 1,
              'item_id': 1,
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'owner': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'type': 0,
              'purpose': '読みたかったから。',
              'due_date': '2019/07/30 23:30:00',
              'created_at': '2019/07/28 22:00:00',
              'updated_at': '2019/07/28 22:00:00'
            }
          ],
          'tags': [
            {
              'id': 1,
              'name': '小説'
            }
          ],
          'like_counts': 1,
          'img_url': 'https://i0.wp.com/sakidorico.s3.amazonaws.com/wp/wp-content/uploads/2019/02/5c7758a49537b.jpg?fit=750%2C750&ssl=1',
          'created_at': '2019/07/28 22:00:00',
          'updated_at': '2019/07/28 22:00:00'
        },
        {
          'id': 2,
          'name': '小説　天気の子',
          'code': 9784041026403,
          'type': 1,
          'owners': [
            {
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': true
            },
            {
              'user': {
                'id': 2,
                'name': 'ryoha',
                'displayName': 'りょは',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': false
            }
          ],
          'latest_logs': [
            {
              'id': 1,
              'item_id': 1,
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'owner': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'type': 0,
              'purpose': '読みたかったから。',
              'due_date': '2019/07/30 23:30:00',
              'created_at': '2019/07/28 22:00:00',
              'updated_at': '2019/07/28 22:00:00'
            }
          ],
          'tags': [
            {
              'id': 1,
              'name': '小説'
            }
          ],
          'like_counts': 114,
          'img_url': 'https://cover.openbd.jp/9784041026403.jpg',
          'created_at': '2019/07/28 22:00:00',
          'updated_at': '2019/07/28 22:00:00'
        },
        {
          'id': 3,
          'name': '天気の子　天気の子　天気の子　天気の子　天気の子',
          'code': 9784041026403,
          'type': 1,
          'owners': [
            {
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': true
            },
            {
              'user': {
                'id': 2,
                'name': 'ryoha',
                'displayName': 'りょは',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': false
            },
            {
              'user': {
                'id': 3,
                'name': 'series2',
                'displayName': 'series2',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': false
              },
              'rentalable': false
            }
          ],
          'latest_logs': [
            {
              'id': 1,
              'item_id': 1,
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'owner': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'type': 0,
              'purpose': '読みたかったから。',
              'due_date': '2019/07/30 23:30:00',
              'created_at': '2019/07/28 22:00:00',
              'updated_at': '2019/07/28 22:00:00'
            }
          ],
          'tags': [
            {
              'id': 1,
              'name': '小説'
            }
          ],
          'img_url': 'https://q.trap.jp/api/1.0/files/f20a1393-4edf-4bcd-993e-9564d2f72293/thumbnail',
          'created_at': '2019/07/28 22:00:00',
          'updated_at': '2019/07/28 22:00:00'
        },
        {
          'id': 4,
          'name': '天気の子',
          'code': 9784041026403,
          'type': 1,
          'owners': [
            {
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': true
            },
            {
              'user': {
                'id': 2,
                'name': 'ryoha',
                'displayName': 'りょは',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'rentalable': false
            },
            {
              'user': {
                'id': 3,
                'name': 'series2',
                'displayName': 'series2',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': false
              },
              'rentalable': false
            },
            {
              'user': {
                'id': 4,
                'name': 'Adwaver_4157',
                'displayName': 'Waver',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': false
              },
              'rentalable': false
            }
          ],
          'latest_logs': [
            {
              'id': 1,
              'item_id': 1,
              'user': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'owner': {
                'id': 1,
                'name': 'nagatech',
                'displayName': 'ながてち',
                'iconFileId': 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
                'admin': true
              },
              'type': 0,
              'purpose': '読みたかったから。',
              'due_date': '2019/07/30 23:30:00',
              'created_at': '2019/07/28 22:00:00',
              'updated_at': '2019/07/28 22:00:00'
            }
          ],
          'tags': [
            {
              'id': 1,
              'name': '小説'
            }
          ],
          'img_url': 'https://q.trap.jp/api/1.0/files/f3aefd3d-8205-46ff-a121-ef19ea85517a/thumbnail',
          'created_at': '2019/07/28 22:00:00',
          'updated_at': '2019/07/28 22:00:00'
        }
      ]
    }
  },
  mounted () {
    // axios
    //   .get(`/api/items`)
    //   .then(res => (this.data = res.data))
    //   .catch(e => { alert(e) })
    // 本番ではaxios.getでマウントしてsampleDataを消してください
    this.data = this.sampleData
  },
  methods: {
    checkRentalable (ownerID) {
      // いい感じにしてください。同じownerが複数いるときのロジックがわかりませんでした
      // 貸し出し可ならfalseを返し不可なら'ryohaが借りてます'みたいなのを返すと思ってます
      return false
    },
    like () {
      // axios.post(/likes)みたいな感じ？
    },
    clickAddOwner () {
      window.open('/register_owner_form', 'newwindow', 'width=400,height=300')
      // window.open('/items/' + this.data.id + '/owner/new', 'newwindow', 'width=400,height=300')
    },
    clickRental () {
      window.open('/rental_form', 'newwindow', 'width=400,height=300')
      // window.open('/items/' + this.data.id + '/rental', 'newwindow', 'width=400,height=300')
    }
  }
}
</script>
