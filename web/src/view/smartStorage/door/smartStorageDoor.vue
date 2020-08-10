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
                     type="primary">新增门禁</el-button>
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
      <el-table-column type="selection"
                       width="55"></el-table-column>
      <el-table-column label="日期"
                       width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="门禁ID"
                       prop="doorId"
                       width="120"></el-table-column>

      <el-table-column label="门禁名称"
                       prop="doorName"
                       width="120"></el-table-column>

      <el-table-column label="门禁状态"
                       width="120">
        <template slot-scope="scope">{{scope.row.doorStatus|filterStatus}}</template>

      </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="updateSmartStorageDoor(scope.row)"
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
                         @click="deleteSmartStorageDoor(scope.row)">确定</el-button>
            </div>
            <el-button type="danger"
                       icon="el-icon-delete"
                       size="mini"
                       slot="reference">删除</el-button>
          </el-popover>
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
      <el-form ref="elForm"
               :model="formData"
               size="medium"
               label-width="100px">
        <el-form-item label="门禁名称"
                      prop="doorName">
          <el-input v-model="formData.doorName"
                    placeholder="请输入门禁名称"
                    clearable
                    :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="门禁状态"
                      prop="doorStatus">
          <el-radio-group v-model="formData.doorStatus"
                          size="medium">
            <el-radio v-for="(item, index) in doorStatusOptions"
                      :key="index"
                      :label="item.value"
                      :disabled="item.disabled">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
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
  createSmartStorageDoor,
  deleteSmartStorageDoor,
  deleteSmartStorageDoorByIds,
  updateSmartStorageDoor,
  findSmartStorageDoor,
  getSmartStorageDoorList,
} from '@/api/smartStorageDoor' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'SmartStorageDoor',
  mixins: [infoList],
  data() {
    return {
      listApi: getSmartStorageDoorList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        doorId: null,
        doorName: null,
        doorStatus: null,
      },
      doorStatusOptions: [
        {
          label: '关闭',
          value: 0,
        },
        {
          label: '开启',
          value: 1,
        },
      ],
    }
  },
  filters: {
    filterStatus: function (status) {
      switch (status) {
        case 0:
          return '关闭'
          break
        case 1:
          return '开启'
          break
        default:
          return '未知'
          break
      }
    },
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
      const res = await deleteSmartStorageDoorByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSmartStorageDoor(row) {
      const res = await findSmartStorageDoor({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.resmartStorageDoor
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        doorId: null,
        doorName: null,
        doorStatus: null,
      }
    },
    async deleteSmartStorageDoor(row) {
      this.visible = false
      const res = await deleteSmartStorageDoor({ ID: row.ID })
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
          res = await createSmartStorageDoor(this.formData)
          break
        case 'update':
          res = await updateSmartStorageDoor(this.formData)
          break
        default:
          res = await createSmartStorageDoor(this.formData)
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