package service

import (
    "gin_generator/internal/core"
    "gin_generator/model"
    "gin_generator/model/dto"
)


// GenTableService 代码生成业务 Service 层
type GenTableService struct {
}

// InsertGenTable 新增代码生成业务
func (this *GenTableService) InsertGenTable(genTable *model.GenTable) (res any, err error) {

    return genTable.InsertGenTable(core.GLOBAL_DB)
}

// UpdateGenTable 修改代码生成业务
func (this *GenTableService) UpdateGenTable(genTable *model.GenTable) (res any, err error) {

    return genTable.UpdateGenTableByTableId(core.GLOBAL_DB)
}

// DeleteGenTable 删除代码生成业务
func (this *GenTableService) DeleteGenTable(tableIdList []int64) (res any, err error) {

    return (&model.GenTable{}).DeleteGenTable(core.GLOBAL_DB, tableIdList)
}

// GetGenTableByTableId 获取代码生成业务业务详细信息
func (this *GenTableService) GetGenTableByTableId(tableId int64) (res any, err error) {

    genTable := model.GenTable{}
    err = (&genTable).GetGenTableByTableId(core.GLOBAL_DB, tableId)
    if err != nil {
        return nil, err
    }

    return genTable, nil
}

// GetGenTableList 查询代码生成业务业务列表
func (this *GenTableService) GetGenTableList(genTableDto *dto.GenTableDto) (res any, err error) {

    genTable, err := genTableDto.DtoToModel()
    genTableList, err := genTable.FindGenTableList(core.GLOBAL_DB, genTableDto.SatrtTime, genTableDto.EndTime)
    if err != nil {
        return nil, err
    }

    return genTableList, nil
}

// GetGenTablePageList 分页查询代码生成业务业务列表
func (this *GenTableService) GetGenTablePageList(genTableDto *dto.GenTableDto) (res any, err error) {

    genTable, err := genTableDto.DtoToModel()
    genTableList, total, err := genTable.FindGenTablePageList(core.GLOBAL_DB, genTableDto.SatrtTime, genTableDto.EndTime, genTableDto.PageSize, genTableDto.PageNum)
    if err != nil {
        return nil, err
    }

    return BuildPageData[model.GenTable](genTableList, total, genTableDto.PageNum, genTableDto.PageSize), nil
}

// ExportGenTable 导出代码生成业务业务列表
func (this *GenTableService) ExportGenTable(genTableDto *dto.GenTableDto) (res any, err error) {

    genTable, err := genTableDto.DtoToModel()
    genTable.FindGenTablePageList(core.GLOBAL_DB, genTableDto.SatrtTime, genTableDto.EndTime, 1, 10000)
    // 实现导出 ...

    return nil, nil
}

func BuildPageData[T any](dataList []T, total int64, pageNum, pageSize int) map[string]any {
    if pageSize <= 0 {
        pageSize = 10 // 默认每页 10 条，避免除零
    }
    if pageNum <= 0 {
        pageNum = 1 // 默认第一页
    }

    pages := int((total + int64(pageSize) - 1) / int64(pageSize))
    size := len(dataList)

    startRow := 0
    endRow := 0
    if size > 0 {
        startRow = (pageNum-1)*pageSize + 1
        endRow = startRow + size - 1
    }

    prePage := 0
    if pageNum > 1 {
        prePage = pageNum - 1
    }

    nextPage := 0
    if pageNum < pages {
        nextPage = pageNum + 1
    }

    isFirstPage := pageNum == 1
    isLastPage := pageNum == pages || pages == 0
    hasPreviousPage := pageNum > 1
    hasNextPage := pageNum < pages

    navigatePages := 10
    navigatepageNums := []int{}
    startNav := pageNum - navigatePages/2
    if startNav < 1 {
        startNav = 1
    }
    endNav := startNav + navigatePages - 1
    if endNav > pages {
        endNav = pages
    }

    for i := startNav; i <= endNav; i++ {
        navigatepageNums = append(navigatepageNums, i)
    }

    navigateFirstPage := 0
    if len(navigatepageNums) > 0 {
        navigateFirstPage = navigatepageNums[0]
    }

    navigateLastPage := 0
    if len(navigatepageNums) > 0 {
        navigateLastPage = navigatepageNums[len(navigatepageNums)-1]
    }

    return map[string]any{
        "total":            total,
        "list":             dataList,
        "pageNum":          pageNum,
        "pageSize":         pageSize,
        "size":             size,
        "startRow":         startRow,
        "endRow":           endRow,
        "pages":            pages,
        "prePage":          prePage,
        "nextPage":         nextPage,
        "isFirstPage":      isFirstPage,
        "isLastPage":       isLastPage,
        "hasPreviousPage":  hasPreviousPage,
        "hasNextPage":      hasNextPage,
        "navigatePages":    navigatePages,
        "navigatepageNums": navigatepageNums,
        "navigateFirstPage": navigateFirstPage,
        "navigateLastPage":  navigateLastPage,
    }
}
