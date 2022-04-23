package top.cyvan.oj.geekbang.week2.recursion;

import top.cyvan.oj.CommonUtil;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;

/**
 * <a href="https://leetcode-cn.com/problems/permutations-ii/">https://leetcode-cn.com/problems/permutations-ii/</a>
 */
public class PermuteUnique {

    public static void main(String[] args) {
        int[] nums = {1, 1, 2};

        PermuteUnique executor = new PermuteUnique();
        List<List<Integer>> lists = executor.permuteUnique(nums);

        lists.forEach(CommonUtil::printIntegerList);
    }

    private List<List<Integer>> permuteUnique(int[] nums) {
        // 1, 2, 3
        // 第一位 1，2，3 选一个        [1]
        // 第二位 2， 3 选一个   [1,2]      [1,3]
        // 第三位 3 选一个   [1,2,3]    第三位 2 选一个 [1,3,2]
        // 维护一个是否选中的数组（map），每次从未选中的里头选一个加入进来
        HashSet<List<Integer>> resSet = new HashSet<>();
        List<Integer> walked = new LinkedList<>();
        // 题意是数字不重复，但是 selected 比 walked.contains 效率高
        boolean[] selected = new boolean[nums.length];

        recur(resSet, walked, selected, nums);

        return new ArrayList<>(resSet);
    }

    private void recur(HashSet<List<Integer>> resSet, List<Integer> walked, boolean[] selected, int[] nums) {
        // 退出条件
        if (walked.size() == nums.length) {
            resSet.add(new ArrayList<>(walked));
            return;
        }

        // 循环遍历 nums，找 !selected 的加入，然后回溯
        for (int i = 0; i < nums.length; ++i) {
            if (!selected[i]) {
                // 如果第 i 个没访问过
                // 排这个元素，看下一个
                selected[i] = true;
                walked.add(nums[i]);
                recur(resSet, walked, selected, nums);

                // 还原
                selected[i] = false;
                walked.remove(walked.size() - 1);
            }
        }
    }

}
