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
                     type="primary">新增配额表</el-button>
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

      <el-table-column label="部门名称"
                       prop="SysAuthority.authorityName"
                       width="120"></el-table-column>

      <el-table-column label="货品名称"
                       prop="SmartStorageProduct.productName"
                       width="120"></el-table-column>

      <el-table-column label="月配额"
                       prop="quotaMonth"
                       width="120"></el-table-column>

      <el-table-column label="月已经领用"
                       prop="usedMonth"
                       width="120"></el-table-column>

      <el-table-column label="年配额"
                       prop="quotaYear"
                       width="120"></el-table-column>

      <el-table-column label="年已经领用"
                       prop="usedYear"
                       width="120"></el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="updateSmartStorageQuota(scope.row)"
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
                         @click="deleteSmartStorageQuota(scope.row)">确定</el-button>
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
               :rules="rules"
               size="medium"
               label-width="100px">
        <el-form-item label="部门名"
                      prop="authorityId">
          <el-select v-model="formData.authorityId"
                     placeholder="请选择部门名"
                     clearable
                     :style="{width: '100%'}">
            <el-option v-for="(item, index) in authorityIdOptions"
                       :key="index"
                       :label="item.label"
                       :value="item.value"
                       :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="货品"
                      prop="productId">
          <el-select v-model="formData.productId"
                     placeholder="请选择货品"
                     clearable
                     :style="{width: '100%'}">
            <el-option v-for="(item, index) in productIdOptions"
                       :key="index"
                       :label="item.label"
                       :value="item.value"
                       :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="月配额"
                      prop="quotaMonth">
          <el-input-number v-model="formData.quotaMonth"
                           placeholder="月配额"></el-input-number>
        </el-form-item>
        <el-form-item label="月已经领用"
                      prop="usedMonth">
          <el-input-number v-model="formData.usedMonth"
                           placeholder="月已经领用"></el-input-number>
        </el-form-item>
        <el-form-item label="年配额"
                      prop="quotaYear">
          <el-input-number v-model="formData.quotaYear"
                           placeholder="年配额"></el-input-number>
        </el-form-item>
        <el-form-item label="年已经领用"
                      prop="usedYear">
          <el-input-number v-model="formData.usedYear"
                           placeholder="年已经领用"></el-input-number>
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
import { getSmartStorageProductList } from '@/api/smartStorageProduct'
import { getAuthorityList } from '@/api/authority'
import {
  createSmartStorageQuota,
  deleteSmartStorageQuota,
  deleteSmartStorageQuotaByIds,
  updateSmartStorageQuota,
  findSmartStorageQuota,
  getSmartStorageQuotaList,
} from '@/api/smartStorageQuota' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'SmartStorageQuota',
  mixins: [infoList],
  data() {
    return {
      listApi: getSmartStorageQuotaList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      rules: {
        authorityId: [
          {
            required: true,
            message: '请选择部门名',
            trigger: 'change',
          },
        ],
        productId: [
          {
            required: true,
            message: '请选择货品',
            trigger: 'change',
          },
        ],
        quotaMonth: [
          {
            required: true,
            message: '月配额',
            trigger: 'blur',
          },
        ],
        usedMonth: [
          {
            required: true,
            message: '月已经领用',
            trigger: 'blur',
          },
        ],
        quotaYear: [
          {
            required: true,
            message: '年配额',
            trigger: 'blur',
          },
        ],
        usedYear: [
          {
            required: true,
            message: '年已经领用',
            trigger: 'blur',
          },
        ],
      },
      authorityIdOptions: [],
      productIdOptions: [],
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        authorityId: null,
        productId: null,
        quotaMonth: 100,
        usedMonth: 0,
        quotaYear: 1000,
        usedYear: 0,
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
      const res = await deleteSmartStorageQuotaByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSmartStorageQuota(row) {
      const res = await findSmartStorageQuota({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.ressq
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        authorityId: null,
        productId: null,
        quotaMonth: 100,
        usedMonth: 0,
        quotaYear: 1000,
        usedYear: 0,
      }
    },
    async deleteSmartStorageQuota(row) {
      this.visible = false
      const res = await deleteSmartStorageQuota({ ID: row.ID })
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
          res = await createSmartStorageQuota(this.formData)
          break
        case 'update':
          res = await updateSmartStorageQuota(this.formData)
          break
        default:
          res = await createSmartStorageQuota(this.formData)
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
      let res
      res = await getAuthorityList({ page: 1, pageSize: 100 })
      if (res.code == 0) {
        this.authorityIdOptions = []
        res.data.list.forEach((element) => {
          this.authorityIdOptions.push({
            value: element.authorityId,
            label: element.authorityName,
          })
        })
      }
      res = await getSmartStorageProductList({ page: 1, pageSize: 100 })
      if (res.code == 0) {
        this.productIdOptions = []
        res.data.list.forEach((element) => {
          this.productIdOptions.push({
            value: element.productId,
            label: element.productName,
          })
        })
      }
      await this.getTableData()
    },
  },
  created() {
    this.init()
  },
}
</script>

<style>
</style>