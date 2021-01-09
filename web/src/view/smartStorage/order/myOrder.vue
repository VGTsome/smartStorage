<template>
  <div>
    <el-row>
      <el-col :span="24">货物低报警</el-col>
    </el-row>

    <el-table :data="tableDataProduct"
              border
              ref="
              multipleTable"
              stripe
              row-class-name="dangerit"
              style="width: 100%"
              tooltip-effect="dark">

      <el-table-column label="货柜编号"
                       prop="SmartStorageCabinet.cabinetName"
                       width="180"></el-table-column>

      <el-table-column label="货物名称"
                       prop="SmartStorageProduct.productName"
                       width="120"></el-table-column>
      <el-table-column label="货物数量"
                       prop="productNumber"
                       width="120"></el-table-column>

    </el-table>
    <el-pagination :current-page="pageProduct"
                   :page-size="pageSizeProduct"
                   :page-sizes="[10,50,100]"
                   :style="{float:'right',padding:'20px'}"
                   :total="totalProduct"
                   @current-change="handleCurrentChangeProduct"
                   @size-change="handleSizeChangeProduct"
                   layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>
    <el-row>
      <el-col :span="24">我的订单</el-col>
    </el-row>
    <div class="search-term">
      <el-form :inline="true"
               :model="searchInfo"
               class="demo-form-inline">
        <el-form-item label="搜索">
          <el-input v-model="searchInfo.searchTxt"
                    placeholder="输入商品的名称"></el-input>
        </el-form-item>

      </el-form>
    </div>
    <el-table :data="tableData.filter(data => !searchInfo.searchTxt || data.SmartStorageProduct.productName.toLowerCase().includes(searchInfo.searchTxt.toLowerCase()))"
              border
              ref="
              multipleTable"
              stripe
              style="width: 100%"
              tooltip-effect="dark">
      <el-table-column label="订货时间"
                       width="200">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="订单编号"
                       prop="orderId"
                       width="180"></el-table-column>

      <el-table-column label="用户昵称"
                       prop="SysUser.nickName"
                       width="120"></el-table-column>
      <el-table-column label="取货名称"
                       prop="SmartStorageProduct.productName"
                       width="120"></el-table-column>
      <el-table-column label="数量"
                       prop="orderNumber"
                       width="120"></el-table-column>
      <el-table-column label="订单状态"
                       sortable
                       :formatter="formatStatus"
                       prop="orderStatus"
                       width="120"></el-table-column>

    </el-table>
    <el-pagination :current-page="page"
                   :page-size="pageSize"
                   :page-sizes="[10,50,100]"
                   :style="{float:'right',padding:'20px'}"
                   :total="total"
                   @current-change="handleCurrentChange"
                   @size-change="handleSizeChange"
                   layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>
  </div>
</template>
<script>
import {
  createSmartStorageOrder,
  deleteSmartStorageOrder,
  deleteSmartStorageOrderByIds,
  updateSmartStorageOrder,
  findSmartStorageOrder,
  getSmartStorageOrderList,
  getAllOrderList,
  getMyOrderList,
  updateSmartStorageOrderStatus,
} from '@/api/smartStorageOrder'
import { getLowNumberProductList } from '@/api/smartStorageProduct'
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'
export default {
  name: 'orderManage',
  mixins: [infoList],
  data() {
    return {
      listApi: getMyOrderList,
      orderCancelVisible: [],
      tableDataProduct: null,
      searchInfo: { searchTxt: '' },
      searchInfoProduct: {},
      pageProduct: 1,
      pageSizeProduct: 10,
      totalProduct: null,
    }
  },
  computed: {},
  filters: {
    formatDate: function (time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
  },
  methods: {
    handleSizeChangeProduct(val) {
      this.pageSizeProduct = val
      this.getTableData()
    },
    handleCurrentChangeProduct(val) {
      this.pageProduct = val
      this.getTableData()
    },
    async fixOrder(val) {
      let res
      val.orderStatus = 10
      res = await updateSmartStorageOrderStatus(val)
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '更改成功',
        })
      }
    },
    async passOrder(val) {
      let res
      val.orderStatus = 0
      res = await updateSmartStorageOrderStatus(val)
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '更改成功',
        })
      }
    },
    async passAllOrder(val) {
      let res
      val.orderStatus = 999
      res = await updateSmartStorageOrderStatus(val)
      val.orderStatus = 0
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '更改成功',
        })
      }
    },
    formatStatus(row, column) {
      let retstatus = ''
      switch (row.orderStatus) {
        case -1:
          retstatus = '待审核'
          break
        case 0:
          retstatus = '审核完成'
          break
        case 1:
          retstatus = '进入仓库'
          break
        case 2:
          retstatus = '出仓盘货'
          break
        case 10:
          retstatus = '订单完成'
          break
        case 401:
          retstatus = '盘货出错'
          break
        default:
          retstatus = '未完成'
      }
      return retstatus
    },
    async getTableDataProduct(
      page = this.pageProduct,
      pageSize = this.pageSizeProduct
    ) {
      const table = await getLowNumberProductList({
        page,
        pageSize,
        ...this.searchInfoProduct,
      })
      this.tableDataProduct = table.data.list
      this.totalProduct = table.data.total
      this.pageProduct = table.data.page
      this.pageSizeProduct = table.data.pageSize
    },
    async init() {
      await this.getTableData()
      await this.getTableDataProduct()
    },
  },
  created() {
    this.init()
  },
}
</script>
<style>
.el-table .dangerit {
  background: oldlace;
}
</style>