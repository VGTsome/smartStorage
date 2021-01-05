<template>
  <div>

    <el-table :data="tableData"
              @selection-change="handleSelectionChange"
              border
              ref="multipleTable"
              stripe
              style="width: 100%"
              tooltip-effect="dark">

      <el-table-column label="日期"
                       width="180">
        <template slot-scope="scope">{{scope.row.UpdatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="系统状态"
                       prop="comments"></el-table-column>

    </el-table>

  </div>
</template>

<script>
import {
  createSmartStorageSystemStatus,
  deleteSmartStorageSystemStatus,
  deleteSmartStorageSystemStatusByIds,
  updateSmartStorageSystemStatus,
  findSmartStorageSystemStatus,
  getSmartStorageSystemStatusList,
} from '@/api/smartStorageSystemStatus' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'SmartStorageSystemStatus',
  mixins: [infoList],
  data() {
    return {
      listApi: getSmartStorageSystemStatusList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        systemStatus: null,
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
    DDDD() {
      setInterval(
        this.onSubmit, //刷新页面
        6000
      )
    }, //这就是 一分钟
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
      const res = await deleteSmartStorageSystemStatusByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSmartStorageSystemStatus(row) {
      const res = await findSmartStorageSystemStatus({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.ressss
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        systemStatus: null,
      }
    },
    async deleteSmartStorageSystemStatus(row) {
      this.visible = false
      const res = await deleteSmartStorageSystemStatus({ ID: row.ID })
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
          res = await createSmartStorageSystemStatus(this.formData)
          break
        case 'update':
          res = await updateSmartStorageSystemStatus(this.formData)
          break
        default:
          res = await createSmartStorageSystemStatus(this.formData)
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
    this.DDDD()
  },
}
</script>

<style>
</style>