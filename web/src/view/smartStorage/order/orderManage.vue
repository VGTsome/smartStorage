<template>
  <div>
    <el-table :data="tableData"
              border
              ref="multipleTable"
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
      <el-table-column label="操作"
                       width="120">
        <template slot-scope="scope">

          <el-popover placement="top"
                      v-model="orderCancelVisible[scope.$index]"
                      width="160">
            <p>确定取消吗？</p>
            <div style="text-align: right; margin: 0">

              <el-button @click="cancelOrder(scope.row)"
                         size="mini"
                         type="primary">确定</el-button>
            </div>
            <el-button slot="reference"
                       :disabled="scope.row.orderStatus!='0'"
                       size="mini"
                       type="danger">强制通过</el-button>
          </el-popover>

        </template>
      </el-table-column>
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
  updateSmartStorageOrderStatus,
} from '@/api/smartStorageOrder'
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'
export default {
  name: 'orderManage',
  mixins: [infoList],
  data() {
    return {
      listApi: getAllOrderList,
      orderCancelVisible: [],
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
    async cancelOrder(val) {
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
    formatStatus(row, column) {
      let retstatus = ''
      switch (row.orderStatus) {
        case 0:
          retstatus = '未完成'
          break
        case 10:
          retstatus = '已完成'
          break
        default:
          retstatus = '未完成'
      }
      return retstatus
    },
    async init() {
      await this.getTableData()
    },
  },
  created() {
    this.init()
  },
}
</script>
