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
                     type="primary">新增货柜产品</el-button>
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
        <el-form-item>
          选择更新货柜
          <el-select v-model="shelf.shelfID"
                     placeholder="选择更新货柜号">
            <el-option value=1></el-option>
            <el-option value=2></el-option>
            <el-option value=3></el-option>
            <el-option value=4></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="updatePrepare"
                     type="primary">全部预备更新</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="updateAll"
                     type="primary">全部更新</el-button>
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

      <el-table-column label="货柜ID"
                       prop="cabinetId"
                       width="120"></el-table-column>
      <el-table-column label="货柜名称"
                       prop="SmartStorageCabinet.cabinetName"
                       width="120"></el-table-column>

      <el-table-column label="产品ID"
                       prop="productId"
                       width="120"></el-table-column>
      <el-table-column label="产品名称"
                       prop="SmartStorageProduct.productName"
                       width="120"></el-table-column>

      <el-table-column label="单个重量(g)"
                       prop="weight"
                       width="120"></el-table-column>
      <el-table-column label="总数量"
                       prop="productNumber"
                       width="120"></el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="updateCabinetProduct(scope.row)"
                     size="small"
                     type="primary">配置</el-button>
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
                         @click="deleteCabinetProduct(scope.row)">确定</el-button>
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
        <el-form-item label="货柜"
                      prop="cabinetId">
          <el-select v-model="formData.cabinetId"
                     placeholder="请选择货柜"
                     clearable
                     :style="{width: '100%'}">
            <el-option v-for="(item, index) in cabinetIdOptions"
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
        <el-form-item label="单个重量(g)"
                      prop="weight">
          <el-input-number v-model="formData.weight"
                           placeholder="单个重量(g)"
                           disabled="1"
                           :min="0"
                           :precision='3'></el-input-number>
        </el-form-item>
        <el-form-item label="总数量"
                      prop="productNumber">
          <el-input-number v-model="formData.productNumber"
                           :min="0"
                           placeholder="总数量"></el-input-number>
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
import { getSmartStorageCabinetList } from '@/api/smartStorageCabinet'
import {
  createCabinetProduct,
  deleteCabinetProduct,
  deleteCabinetProductByIds,
  updateCabinetProduct,
  findCabinetProduct,
  getCabinetProductList,
  prepareShelf,
  updateShelf,
} from '@/api/cabinetProduct' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'CabinetProduct',
  mixins: [infoList],
  data() {
    return {
      listApi: getCabinetProductList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        cabinetId: null,
        productId: null,
        weight: null,
        productNumber: null,
      },
      shelf: {
        shelfID: 1,
        action: 1,
      },
      rules: {
        cabinetId: [
          {
            required: true,
            message: '请选择货柜',
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
        weight: [
          {
            required: true,
            message: '单个重量(g)',
            trigger: 'blur',
          },
        ],
      },
      cabinetIdOptions: [],
      productIdOptions: [],
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
    updateAll() {
      //this.page = 1
      this.shelf.action = 2
      this.updateShelf(this.shelf)
    },
    updatePrepare() {
      //this.page = 1
      this.shelf.action = 1
      this.prepareShelf(this.shelf)
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
      const res = await deleteCabinetProductByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateCabinetProduct(row) {
      const res = await findCabinetProduct({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.resscp
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        cabinetId: null,
        productId: null,
        status: null,
      }
    },
    async deleteCabinetProduct(row) {
      this.visible = false
      const res = await deleteCabinetProduct({ ID: row.ID })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.getTableData()
      }
    },
    async prepareShelf(data) {
      const res = await prepareShelf(data)
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '准备成功',
        })
      }
    },
    async updateShelf(data) {
      const res = await updateShelf(data)
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '更新成功',
        })
      }
    },
    async enterDialog() {
      let valid2
      await this.$refs['elForm'].validate((valid) => {
        valid2 = valid
        if (!valid) return
      })
      if (valid2) {
        this.formData.weight = this.formData.weight
        let res
        switch (this.type) {
          case 'create':
            res = await createCabinetProduct(this.formData)
            break
          case 'update':
            res = await updateCabinetProduct(this.formData)
            break
          default:
            res = await createCabinetProduct(this.formData)
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
        this.$refs['elForm'].resetFields()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
    async init() {
      let res
      res = await getSmartStorageCabinetList({ page: 1, pageSize: 100 })
      if (res.code == 0) {
        this.cabinetIdOptions = []
        res.data.list.forEach((element) => {
          this.cabinetIdOptions.push({
            value: element.cabinetId,
            label: element.cabinetName,
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