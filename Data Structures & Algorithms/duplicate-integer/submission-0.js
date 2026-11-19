class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        for (const num of nums) {
            if (nums.lastIndexOf(num) > nums.indexOf(num)) {
                return true
            }
        }

        return false
    }
}
