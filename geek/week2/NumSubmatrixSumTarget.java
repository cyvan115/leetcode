package top.cyvan.oj.geekbang.week2.map;

import java.util.HashMap;
import java.util.Map;

/**
 * <a href="https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/">https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/</a>
 */
public class NumSubmatrixSumTarget {

    public static void main(String[] args) {
        int[][] matrix = {
                {0, 1, 0},
                {1, 1, 1},
                {0, 1, 0}
        };
        int target = 0;

        NumSubmatrixSumTarget executor = new NumSubmatrixSumTarget();
        System.out.println(executor.numSubmatrixSumTarget(matrix, target));
    }

    private int numSubmatrixSumTarget(int[][] matrix, int target) {
        // 通过定义上下界，得到最多有多少行元素
        int res = 0;
        int rows = matrix.length;
        int cols = matrix[0].length;

        for (int i = 0; i < rows; ++i) {
            int[] colSum = new int[cols];
            for (int j = i; j < rows; ++j) {
                // 算出每一列的值，统一抽象为一维的数组进行处理
                for (int col = 0; col < cols; ++col) {
                    // 依赖上一个抽象行计算新的抽象行
                    // 每次加上当前 j 行的第 col 个元素即可
                    colSum[col] += matrix[j][col];
                }

                res += findSubSumInArray(colSum, target);
            }
        }

        return res;
    }

    private int findSubSumInArray(int[] nums, int target) {
        // 前缀和数组
        // 如果第 i 个元素到 j 个元素之前和满足和为 target
        // nums[j] - nums[i - 1] = target
        // 所以满足 nums[i - 1] = nums[j] - target
        // 判断 nums[j] - target 是否出现过即可
        Map<Integer, Integer> sumMap = new HashMap<>();
        sumMap.put(0, 1);
        int preSum = 0;
        int cnt = 0;
        for (int i = 0; i < nums.length; ++i) {
            preSum += nums[i];
            if (sumMap.containsKey(preSum - target)) {
                cnt += sumMap.get(preSum - target);
            }

            // 如果有两个前缀和 = 同一个数，表示 这俩之间的元素和为 0
            sumMap.put(preSum, sumMap.getOrDefault(preSum, 0) + 1);
        }

        return cnt;
    }

}
