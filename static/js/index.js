new Vue({
    el: "#root",
    data() {
        return {
            curr_page:1,
            page_num:10,
            numCount:1,
            pageCount:10,
            tableData: [],
            commentData:[],
            btnFlag: false,
            showAddDialog: false,
            base_url: location.origin,
            // base_url: "http://localhost:8888",
        }
    },
    methods: {
        getData() {
            fly.get(this.base_url + '/getList?curr_page=' + this.curr_page + '&page_num=' + this.page_num).then(res => {
                var res = res.data
                if (res.code == 200) {
                    this.tableData = res.data.list
                    this.numCount = res.data.numCount
                    this.pageCount = res.data.pageCount
                } else {
                    this.$message({
                        showClose: true,
                        message: res.msg,
                        type: 'error'
                    });
                }
            })
        },
        getComment(){
            fly.get(this.base_url + '/getComment').then(res => {
                var res = res.data
                if (res.code == 200) {
                    this.commentData = res.data
                } else {
                    this.$message({
                        showClose: true,
                        message: res.msg,
                        type: 'error'
                    });
                }
            })
        },
        showDialog(){
            this.showAddDialog = true
            this.getComment()
        },
        handleSizeChange(e){
            console.log(e)
            this.page_num = e
            this.getData()
        },
        handleCurrentChange(e){
            console.log(e)
            this.curr_page = e
            this.getData()
        },
        // 点击图片回到顶部方法，加计时器是为了过渡顺滑
        backTop () {
            let that = this
            let timer = setInterval(() => {
                let ispeed = Math.floor(-that.scrollTop / 5)
                document.documentElement.scrollTop = document.body.scrollTop = that.scrollTop + ispeed
                if (that.scrollTop === 0) {
                    clearInterval(timer)
                }
            }, 16)
        },
        // 为了计算距离顶部的高度，当高度大于60显示回顶部图标，小于60则隐藏
        scrollToTop () {
            let that = this
            let scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop
            that.scrollTop = scrollTop
            if (that.scrollTop > 60) {
                that.btnFlag = true
            } else {
                that.btnFlag = false
            }
        }
    },
    mounted(){
        this.getData()
        window.addEventListener('scroll', this.scrollToTop)
    },
    destroyed () {
        window.removeEventListener('scroll', this.scrollToTop)
    }
})

//
// new Vue({
//     el: "#dialog",
//     data() {
//         return {
//             commentData: [
//                 {
//                     "head": "http://img1.qunarzz.com/sight/p0/1707/f4/f44d291b76de7b14a3.img.jpg_710x360_0df65faf.jpg",
//                     "nick_name": "aaa",
//                     "comment": "好玩"
//                 },
//                 {
//                     "head": "http://img1.qunarzz.com/sight/p0/1707/f4/f44d291b76de7b14a3.img.jpg_710x360_0df65faf.jpg",
//                     "nick_name": "aaa",
//                     "comment": "好玩"
//                 },
//                 {
//                     "head": "http://img1.qunarzz.com/sight/p0/1707/f4/f44d291b76de7b14a3.img.jpg_710x360_0df65faf.jpg",
//                     "nick_name": "aaa",
//                     "comment": "好玩"
//                 }
//             ]
//         }
//     }
// })