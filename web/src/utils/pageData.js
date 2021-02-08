export default {
    data() {
        return {
            page: 1,
            total: 10,
            pageSize: 10,
            tableData: [],
            searchInfo: {}
        }
    },
    methods: {
        getChangePageSizeData(size) {
            this.pageSize = size
            this.getTableData()
        },
        getPageData(page) {
            this.page = page
            this.getTableData()
        },
        async getTableData(page = this.page, pageSize = this.pageSize) {
            const table = await this.listApi({page: page, pageSize: pageSize, ...this.searchInfo})
            //接口成功返回
            if (table.code == 200) {
                if (table.page && table.page > 0) {
                    //有分页数据
                    this.tableData = table.data.list
                    this.total = table.data.total
                    this.page = table.data.page
                    this.pageSize = table.data.page_size
                } else {
                    //没有分页数据
                    this.tableData = table.data
                }
            }
        }
    }
}