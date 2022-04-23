package top.cyvan.oj.geekbang.week2.devide;

public class MergeKLists {

    public static void main(String[] args) {
        MergeKLists executor = new MergeKLists();
        executor.mergeKLists(null);
    }

    private ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) {
            return null;
        }

        return mergeKListWithRange(lists, 0, lists.length - 1);
    }

    private ListNode mergeKListWithRange(ListNode[] lists, int begin, int end) {
        if (begin == end) {
            return lists[begin];
        }

        if (end - begin == 1) {
            return mergeTwoList(lists[begin], lists[end]);
        }

        return mergeTwoList(
                mergeKListWithRange(lists, begin, (begin + end) / 2),
                mergeKListWithRange(lists, (begin + end) / 2 + 1, end)
        );
    }

    private ListNode mergeTwoList(ListNode list1, ListNode list2) {
        ListNode head = new ListNode();

        ListNode cur = head;
        while (list1 != null || list2 != null) {
            if (list1 == null) {
                cur.next = list2;
                break;
            } else if (list2 == null) {
                cur.next = list1;
                break;
            } else {
                if (list1.val <= list2.val) {
                    cur.next = list1;
                    cur = cur.next;
                    list1 = list1.next;
                } else {
                    cur.next = list2;
                    cur = cur.next;
                    list2 = list2.next;
                }
            }
        }

        return head.next;
    }

    static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) {
            this.val = val;
        }
        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }

}
