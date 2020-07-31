<template>
  <div>
    <div class="search-term">
      <el-form :inline="true"
               :model="searchInfo"
               class="demo-form-inline">
        <el-form-item>
          <el-button @blur="onSubmit"
                     type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog"
                     type="primary">购物车</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table :data="tableData"
              @selection-change="handleSelectionChange"
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

      <el-table-column label="物料编码"
                       prop="productId"
                       width="120"></el-table-column>

      <el-table-column label="按钮组">
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
                   :page-sizes="[10, 30, 50, 100]"
                   :style="{float:'right',padding:'20px'}"
                   :total="total"
                   @current-change="handleCurrentChange"
                   @size-change="handleSizeChange"
                   layout="total, sizes, prev, pager, next, jumper"></el-pagination>

    <el-dialog :before-close="closeDialog"
               :visible.sync="dialogFormVisible"
               title="弹窗操作">
      <el-table :data="cartData"
                @selection-change="handleSelectionChange"
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
            <el-button @click="enterDialog"
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
  getSmartStorageProductList,
} from '@/api/smartStorageProduct' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'SmartStorageOrder',
  mixins: [infoList],
  data() {
    return {
      listApi: getSmartStorageProductList,
      dialogFormVisible: false,
      enterVisible: false,
      visible: false,
      type: '',
      cartData: [],
      formData: [],
    }
  },
  filters: {},
  methods: {
    updateCart(val) {
      this.cartData.forEach((item, $index) => {
        if (item.ID == val.ID) {
          this.cartData.pop($index)
        }
      })
      this.cartData.push(val)
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID)
        })
      const res = await deleteSmartStorageOrderByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSmartStorageOrder(row) {
      const res = await findSmartStorageOrder({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.resmartStorageOrder
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = []
    },
    async deleteSmartStorageOrder(row) {
      this.visible = false
      const res = await deleteSmartStorageOrder({ ID: row.ID })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.getTableData()
      }
    },
    async enterDialog() {
      const fdata = []
      this.cartData.forEach((item) => {
        fdata.push({
          productId: item.productId,
          orderNumber: item.orderNumber,
        })
      })
      let res
      switch (this.type) {
        case 'create':
          res = await createSmartStorageOrder({ fdata })
          break
        case 'update':
          res = await updateSmartStorageOrder({ fdata })
          break
        default:
          res = await createSmartStorageOrder({ fdata })
          break
      }
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
        })
        this.enterVisible = false
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      //this.type = 'create'
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