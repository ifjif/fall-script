/*
 * slice:
 * [start : end : step : max]
 *
 * step 不能为0, 默认为1
 * 当step 不是默认值，而是指定值，如果max存在，panic
 * 当step 是默认值，而slice是string, 如果max存在,panic
 *
 * step > 0 不为默认值 时
 * 0 <= start < end <= max <= rawMax (end, max, rawMax 为最大不可到达的索引位置)
 * start 省略 默认 0
 * end   省略 默认 len(arr)
 *
 * step < 0 时（会 copy）
 * len(arr) > start > end >= -1
 * start, end 可为 负数,需要保证
 *    len(arr) > len(arr) + start > len(arr) + end
 * start 省略 默认为 len(arr) - 1
 * end   省略 默认为 -1
 *
 *
 * sliceExpr {
 *  start node
 *  end node
 *  step node
 *  max node
 * }
 *
 * a[sliceExpr]
 *
 * let arr = [1,2,3,4]
 * 
 * arr[1]
 * arr[1:]
 * arr[1:2]
 * arr[1:2:3]
 * arr[1:2:3:4]
 * 
 * arr[:1]
 * arr[:1:2]
 * arr[:1:2:3]
 * 
 * arr[::1]
 * arr[::1:2]
 * 
 * arr[:::1]
 * 
 * arr[:]
 * arr[::]
 * arr[:::]
 *
 */

