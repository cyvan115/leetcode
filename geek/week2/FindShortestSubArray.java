package top.cyvan.oj.geekbang.week2.map;

import java.util.HashMap;
import java.util.Map;

/**
 * <a href="https://leetcode-cn.com/problems/degree-of-an-array/">https://leetcode-cn.com/problems/degree-of-an-array/</a>
 */
public class FindShortestSubArray {

    public static void main(String[] args) {
        int[] nums = {1,2,2,3,1};
//        int[] nums = {1,2,2,3,1,4,2};
//        int[] nums = {1};

        FindShortestSubArray executor = new FindShortestSubArray();

        System.out.println(executor.findShortestSubArray(nums));
    }

    private int findShortestSubArray(int[] nums) {
        // 第一次遍历找出度，并且计算每个元素出现的首、末位置
        Map<Integer, Integer> cnt = new HashMap<>();
        Map<Integer, Integer> head = new HashMap<>();
        Map<Integer, Integer> tail = new HashMap<>();

        int maxCount = 0;

        for (int i = 0; i < nums.length; ++i) {
            cnt.put(nums[i], cnt.getOrDefault(nums[i], 0) + 1);

            maxCount = Math.max(cnt.get(nums[i]), maxCount);

            tail.put(nums[i], i);

            if (!head.containsKey(nums[i])) {
                head.put(nums[i], i);
            }
        }

        int len = nums.length;

        // 根据度计算 末 - 首 的最小值
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            if (e.getValue() == maxCount) {
                len = Math.min(tail.get(e.getKey()) - head.get(e.getKey()) + 1, len);
            }
        }

        return len;
    }



}
