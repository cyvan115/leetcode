package top.cyvan.oj.geekbang.week1.array;

import top.cyvan.oj.CommonUtil;

public class PlusOne {


    public static void main(String[] args) {
        PlusOne executor = new PlusOne();

//        int[] digits = new int[]{1, 2, 3};

//        int[] digits = new int[]{4, 3, 2, 1};

//        int[] digits = new int[]{0};

        int[] digits = new int[]{9, 9, 9, 9};

        int[] ints = executor.plusOne(digits);

        CommonUtil.printIntArray(ints);
    }

    private int[] plusOne(int[] digits) {
        int extra = 1;
        for (int i = digits.length - 1; i >= 0; --i) {
            int sum = digits[i] + extra;

            if (sum == 10) {
                digits[i] = 0;
                extra = 1;
            } else {
                digits[i] = sum;
                extra = 0;

                break;
            }
        }

        if (extra == 1) {
            int[] res = new int[digits.length + 1];
            res[0] = 1;

            for (int i = 0; i < digits.length; ++i) {
                res[i + 1] = digits[i];
            }

            return res;
        }

        return digits;
    }
}
