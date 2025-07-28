//将一个日期字符串（或 ISO 日期）格式化为中文格式的字符串，
export const formatDate = (dateString) => {
    if (!dateString || dateString === '0001-01-01T00:00:00Z') return ''
    const date = new Date(dateString)
    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    })
}

//它会把选择的两个时n间（开始时间和结束时间）格式化后，分别赋值给 form 中的两个字段（默认是 startTime 和 endTime）。
export const newTimeRangeHandler = (form, startField = 'startTime', endField = 'endTime') => {
  return (dates) => {
    if (dates ) {
      form[startField] = dates[0].format('YYYY-MM-DD HH:mm:ss')
      form[endField] = dates[1].format('YYYY-MM-DD HH:mm:ss')
    } else {
      form[startField] = ''
      form[endField] = ''
    }
  }
}

