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
          <el-button @click="openDialog"
                     type="primary">新增货品</el-button>
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
    <el-table :data="tableData.filter(data => !searchInfo.productName || data.productName.toLowerCase().includes(searchInfo.productName.toLowerCase()))"
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

      <el-table-column label="货品编码"
                       prop="productId"
                       width="120"></el-table-column>

      <el-table-column label="货品名称"
                       prop="productName"
                       width="120"></el-table-column>

      <el-table-column label="单位重量（g）"
                       prop="productWeight"
                       width="120"></el-table-column>

      <el-table-column label="货品描述"
                       prop="productDescription"
                       width="120"></el-table-column>

      <el-table-column label="图片"
                       prop="productImgUrl"
                       width="120">
        <template slot-scope="scope"><img :src=scope.row.productNumber /></template></el-table-column>

      <el-table-column label="货品数量"
                       prop="productNumber"
                       width="120">
      </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="updateSmartStorageProduct(scope.row)"
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
                         @click="deleteSmartStorageProduct(scope.row)">确定</el-button>
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
        <el-form-item label="货品编号"
                      prop="productId">
          <el-input v-model="formData.productId"
                    placeholder="请输入货品编号"
                    clearable
                    :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="货品名称"
                      prop="productName">
          <el-input v-model="formData.productName"
                    placeholder="请输入货品名称"
                    clearable
                    :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="货品重量"
                      prop="productWeight">
          <el-input-number v-model="formData.productWeight"
                           placeholder="货品重量"></el-input-number>
        </el-form-item>
        <el-form-item label="货品描述"
                      prop="productDescription">
          <el-input v-model="formData.productDescription"
                    type="textarea"
                    placeholder="请输入货品描述"
                    :autosize="{minRows: 4, maxRows: 4}"
                    :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label-width="0"
                      prop="productImg"
                      required>
          <el-upload ref="productImg"
                     :file-list="productImgfileList"
                     :action="productImgAction"
                     :on-success="updatefilelist"
                     :before-upload="productImgBeforeUpload"
                     list-type="picture"
                     :headers="headinfo"
                     :limit="1">
            <el-button size="small"
                       type="primary"
                       icon="el-icon-upload">点击上传</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="货品数量"
                      prop="filename">
          <el-input-number v-model="formData.productNumber"
                           placeholder="货品数量"></el-input-number>
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
  createSmartStorageProduct,
  deleteSmartStorageProduct,
  deleteSmartStorageProductByIds,
  updateSmartStorageProduct,
  findSmartStorageProduct,
  getSmartStorageProductList,
} from '@/api/smartStorageProduct' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'
import { store } from '@/store/index'
const path = process.env.VUE_APP_BASE_API
const token = store.getters['user/token']
const user = store.getters['user/userInfo']
export default {
  name: 'SmartStorageProduct',

  mixins: [infoList],
  data() {
    return {
      headinfo: {
        'x-token': token,
        'x-user-id': user.ID,
      },
      searchInfo: {
        name: '',
      },
      listApi: getSmartStorageProductList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        productId: null,
        productName: null,
        productWeight: null,
        productDescription: null,
        productImgUrl: null,
        productNumber: null,
      },
      rules: {
        productId: [
          {
            required: true,
            message: '请输入货品编号',
            trigger: 'blur',
          },
        ],
        productName: [
          {
            required: true,
            message: '请输入货品名称',
            trigger: 'blur',
          },
        ],
        productWeight: [
          {
            required: true,
            message: '货品重量',
            trigger: 'blur',
          },
        ],
        productDescription: [
          {
            required: true,
            message: '请输入货品描述',
            trigger: 'blur',
          },
        ],
        productNumber: [],
      },

      productImgAction: path + '/fileUploadAndDownload/upload',
      productImgfileList: [],
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
    updatefilelist(response) {
      this.formData.productImgUrl = response.data.file.url
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
      const res = await deleteSmartStorageProductByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSmartStorageProduct(row) {
      const res = await findSmartStorageProduct({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.resmartStorageProduct
        this.dialogFormVisible = true
      }
    },
    productImgBeforeUpload(file) {
      let isRightSize = file.size / 1024 / 1024 < 5
      if (!isRightSize) {
        this.$message.error('文件大小超过 5MB')
      }
      return isRightSize
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        productId: null,
        productName: null,
        productWeight: null,
        productDescription: null,
        productImgUrl: null,
        productNumber: null,
      }
    },
    async deleteSmartStorageProduct(row) {
      this.visible = false
      const res = await deleteSmartStorageProduct({ ID: row.ID })
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
          res = await createSmartStorageProduct(this.formData)
          break
        case 'update':
          res = await updateSmartStorageProduct(this.formData)
          break
        default:
          res = await createSmartStorageProduct(this.formData)
          break
      }
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
        })
        this.closeDialog()
        this.getTableData()
        this.$refs.productImg.clearFiles()
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