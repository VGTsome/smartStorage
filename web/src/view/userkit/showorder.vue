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
                     type="danger">开门</el-button>
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
            <!--el-button icon="el-icon-delete"
                       size="mini"
                       slot="reference"
                       type="danger">批量删除</！el-button-->
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
      <el-table-column type="selection"
                       width="55"></el-table-column>
      <el-table-column label="订单时间"
                       width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="订单编号"
                       prop="orderId"
                       width="220"></el-table-column>

      <el-table-column label="用户编号"
                       prop="userId"
                       width="120"></el-table-column>

      <el-table-column label="货品名称"
                       prop="productName"
                       width="120"></el-table-column>
      <el-table-column label="货架号"
                       prop="cabinetList"
                       width="120"></el-table-column>

      <el-table-column label="订单数"
                       prop="orderNumber"
                       width="120"></el-table-column>

      <!--el-table-column label="订单状态"
                       prop="orderStatus"
                       width="120"></!--el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateSmartStorageCurrentOrder(scope.row)"
                     size="small"
                     type="primary">变更</el-button>
          <el-popover placement="top"
                      width="160"
                      v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini"
                         type="text"
                         @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary"
                         size="mini"
                         @click="deleteSmartStorageCurrentOrder(scope.row)">确定</el-button>
            </div>
            <el-button type="danger"
                       icon="el-icon-delete"
                       size="mini"
                       slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column-->
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
               title="操作">
      确定已拿齐货品吗？
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
  createSmartStorageCurrentOrder,
  deleteSmartStorageCurrentOrder,
  deleteSmartStorageCurrentOrderByIds,
  updateSmartStorageCurrentOrder,
  findSmartStorageCurrentOrder,
  getSmartStorageCurrentOrderList,
} from '@/api/smartStorageCurrentOrder' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'SmartStorageCurrentOrder',
  mixins: [infoList],
  data() {
    return {
      listApi: getSmartStorageCurrentOrderList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        orderId: null,
        userId: null,
        productId: null,
        productName: null,
        cabinetList: null,
        orderNumber: null,
        orderStatus: null,
      },
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
      const res = await deleteSmartStorageCurrentOrderByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSmartStorageCurrentOrder(row) {
      const res = await findSmartStorageCurrentOrder({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.resmartStorageCurrentOrder
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        orderId: null,
        userId: null,
        productId: null,
        orderNumber: null,
        orderStatus: null,
      }
    },
    async deleteSmartStorageCurrentOrder(row) {
      this.visible = false
      const res = await deleteSmartStorageCurrentOrder({ ID: row.ID })
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
          res = await createSmartStorageCurrentOrder(this.formData)
          break
        case 'update':
          res = await updateSmartStorageCurrentOrder(this.formData)
          break
        default:
          res = await createSmartStorageCurrentOrder(this.formData)
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
  },
  created() {
    this.getTableData()
  },
}
</script>

<style>
</style>