class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        if (nums.length < 2) {
            return false
        }

        return [...new Set(nums)].length !== nums.length
    }
}
