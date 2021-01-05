<template>
  <div>
    <div class="search-term">
      <el-form :inline="true"
               :model="searchInfo"
               class="demo-form-inline">
        <el-form-item label="货品名称">
          <el-input v-model="searchInfo.productName"
                    placeholder="输入搜索的名称"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button @click="openCartDialog"
                     type="primary">购物车 {{this.cartData.length}}</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openOrderDialog"
                     type="primary">查看订单</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table :data="tableData.filter(data => !searchInfo.productName || data.productName.toLowerCase().includes(searchInfo.productName.toLowerCase()))"
              border
              ref="multipleTable"
              stripe
              style="width: 100%"
              tooltip-effect="dark">

      <el-table-column label="图片"
                       width="180">
        <template slot-scope="scope">
          <img width="40"
               v-if="scope.row.productImgUrl!=''"
               :src=scope.row.productImgUrl />
          <img width="40"
               v-else
               src="@/assets/goods.png" />
        </template>
      </el-table-column>

      <el-table-column label="货名"
                       prop="productName"
                       width="120"></el-table-column>

      <el-table-column label="货品描述"
                       prop="productDescription"
                       width="120"></el-table-column>

      <el-table-column label="物料编码"
                       sortable
                       prop="productId"
                       width="120"></el-table-column>

      <el-table-column label="预定数量">

        <template slot-scope="scope">
          <el-input-number placeholder="计数器"
                           v-model=scope.row.orderNumber
                           @change="updateCart(scope.row)"
                           :min=0></el-input-number>
        </template>

      </el-table-column>
    </el-table>

    <el-pagination :current-page="page"
                   :page-size="pageSize"
                   :page-sizes="[100]"
                   :style="{float:'right',padding:'20px'}"
                   :total="total"
                   @current-change="handleCurrentChange"
                   @size-change="handleSizeChange"
                   layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>
    <el-dialog :before-close="closeDialog"
               :visible.sync="orderDiglogVisible"
               title="订单管理">
      <el-table :data="orderData"
                border
                ref="multipleTable"
                stripe
                style="width: 100%"
                tooltip-effect="dark">
        <el-table-column label="图片"
                         width="120">
          <template slot-scope="scope"><img width="40"
                 v-if="scope.row.productImgUrl!=''"
                 :src=scope.row.SmartStorageProduct.productImgUrl /></template>
        </el-table-column>
        <el-table-column label="货名"
                         prop="SmartStorageProduct.productName"
                         width="120"></el-table-column>
        <el-table-column label="货品描述"
                         prop="SmartStorageProduct.productDescription"
                         width="120"></el-table-column>
        <el-table-column label="订货数量"
                         prop="orderNumber"
                         width="120"></el-table-column>
        <el-table-column label="订单状态"
                         prop="orderStatus"
                         :formatter="formatStatus"
                         width="120"></el-table-column>
      </el-table>
      <div class="dialog-footer"
           slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-popover placement="top"
                    v-model="orderCancelVisible"
                    width="160">
          <p>确定取消吗？</p>
          <div style="text-align: right; margin: 0">
            <el-button @click="orderCancelVisible = false"
                       size="mini"
                       type="text">取消</el-button>
            <el-button @click="cancelOrder"
                       size="mini"
                       type="primary">确定</el-button>
          </div>
          <el-button slot="reference"
                     type="danger">取消订单</el-button>
        </el-popover>
      </div>
    </el-dialog>
    <el-dialog :before-close="closeDialog"
               :visible.sync="dialogFormVisible"
               title="购物车">
      <el-table :data="cartData"
                border
                ref="multipleTable"
                stripe
                style="width: 100%"
                tooltip-effect="dark">
        <el-table-column label="图片"
                         width="180">
          <template slot-scope="scope"><img :src=scope.row.productImgUrl /></template>
        </el-table-column>
        <el-table-column label="货名"
                         prop="productName"
                         width="120"></el-table-column>
        <el-table-column label="货品描述"
                         prop="productDescription"
                         width="120"></el-table-column>
        <el-table-column label="订货数量"
                         prop="orderNumber"
                         width="120"></el-table-column>
      </el-table>
      <div class="dialog-footer"
           slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-popover placement="top"
                    v-model="enterVisible"
                    width="160">
          <p>确定提交吗？</p>
          <div style="text-align: right; margin: 0">
            <el-button @click="enterVisible = false"
                       size="mini"
                       type="text">取消</el-button>
            <el-button @click="submitOrder"
                       size="mini"
                       type="primary">确定</el-button>
          </div>
          <el-button slot="reference"
                     type="primary">提交订单</el-button>
        </el-popover>
      </div>
    </el-dialog>
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
} from '@/api/smartStorageOrder' //  此处请自行替换地址
import {
  findSmartStorageProduct,
  getSmartStorageProductValidList,
} from '@/api/smartStorageProduct' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'SmartStorageOrder',
  mixins: [infoList],
  data() {
    return {
      searchInfo: {
        name: '',
      },
      listApi: getSmartStorageProductValidList,
      dialogFormVisible: false,
      orderDiglogVisible: false,
      enterVisible: false,
      orderCancelVisible: false,
      cartData: [],
      orderData: [],
      formData: [],
    }
  },
  filters: {},
  methods: {
    formatStatus(row, column) {
      let retstatus = ''
      switch (row.orderStatus) {
        case -1:
          retstatus = '提交待审核'
          break
        case 1:
          retstatus = '已审核'
          break
        case 10:
          retstatus = '已完成'
          break
        default:
          retstatus = '未完成'
      }
      return retstatus
    },
    updateCart(val) {
      this.cartData.forEach((item, $index) => {
        if (item.ID == val.ID) {
          this.cartData.splice($index, 1)
        }
      })
      if (val.orderNumber > 0) this.cartData.push(val)
    },
    //条件搜索前端看此方法

    async cancelOrder() {
      const ids = []
      this.orderData.map((item) => {
        ids.push(item.ID)
      })
      const res = await deleteSmartStorageOrderByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.closeDialog()
      }
      this.cartData = []
      this.init()
    },
    async openOrderDialog() {
      var res
      res = await getSmartStorageOrderList({ page: 1, pageSize: 100 })
      if (res.code == 0) {
        this.orderData = res.data.list
        this.orderDiglogVisible = true
      }
    },
    closeDialog() {
      this.orderDiglogVisible = false
      this.dialogFormVisible = false
      this.enterVisible = false
      this.orderCancelVisible = false
    },
    async submitOrder() {
      if (this.cartData.length == 0) {
        this.$message({
          type: 'fail',
          message: '购物车无内容',
        })
        return
      }
      const fdata = []
      this.cartData.forEach((item) => {
        fdata.push({
          productId: item.productId,
          orderNumber: item.orderNumber,
        })
      })
      let res

      res = await createSmartStorageOrder({ fdata })

      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
        })

        this.closeDialog()
        this.cartData = []
        this.init()
      }
    },
    openCartDialog() {
      this.dialogFormVisible = true
    },
    async init() {
      await this.getTableData()
      this.tableData.forEach((item) => {
        item.orderNumber = 0
      })
    },
  },
  created() {
    this.init()
  },
}
</script>

<style>
</style>