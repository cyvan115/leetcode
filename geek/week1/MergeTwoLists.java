package top.cyvan.oj.geekbang.week1.linkedlist;

public class MergeTwoLists {

    public static void main(String[] args) {
//        ListNode list1 = new ListNode(1, new ListNode(
//                2, new ListNode(
//                        4)));
//        ListNode list2 = new ListNode(1, new ListNode(
//                3, new ListNode(
//                4)));

//        ListNode list1 = null;
//        ListNode list2 = null;

//        ListNode list1 = null;
//        ListNode list2 = new ListNode(0);

        ListNode list1 = new ListNode(1);
        ListNode list2 = null;

        MergeTwoLists executor = new MergeTwoLists();
        ListNode listNode = executor.mergeTwoList(list1, list2);

        while (listNode != null) {
            System.out.println(listNode.val);

            listNode = listNode.next;
        }
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
