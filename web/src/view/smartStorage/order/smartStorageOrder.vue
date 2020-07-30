<template>
  <div>
    <div class="search-term">
      <el-form :inline="true"
               :model="searchInfo"
               class="demo-form-inline">
        <el-form-item>
          <el-button @click="onSubmit"
                     type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog"
                     type="primary">新增smartStorageOrder表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top"
                      v-model="deleteVisible"
                      width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false"
                         size="mini"
                         type="text">取消</el-button>
              <el-button @click="onDelete"
                         size="mini"
                         type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete"
                       size="mini"
                       slot="reference"
                       type="danger">批量删除</el-button>
          </el-popover>
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

      <el-table-column label="productId字段"
                       prop="productId"
                       width="120"></el-table-column>

      <el-table-column label="orderStatus字段"
                       prop="orderStatus"
                       width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-input-number placeholder="计数器"
                           v-model=scope.row.orderNumber
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
      此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
      <div class="dialog-footer"
           slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog"
                   type="primary">确 定</el-button>
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
      tableData2: [],
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: [
        { orderId: null, userId: null, productId: null, orderStatus: null },
      ],
    }
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    },
  },
  methods: {
    updateCart(val) {
      console.log(val)
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
      this.formData = {
        orderId: null,
        userId: null,
        productId: null,
        orderStatus: null,
      }
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
      let res
      switch (this.type) {
        case 'create':
          res = await createSmartStorageOrder(this.formData)
          break
        case 'update':
          res = await updateSmartStorageOrder(this.formData)
          break
        default:
          res = await createSmartStorageOrder(this.formData)
          break
      }
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
    async init() {
      debugger
      await this.getTableData()

      this.tableData.forEach((item) => {
        item.orderNumber = 2
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