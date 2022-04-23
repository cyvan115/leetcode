package top.cyvan.oj.geekbang.week2.map;

import top.cyvan.oj.CommonUtil;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

/**
 * <a href="https://leetcode-cn.com/problems/subdomain-visit-count/">https://leetcode-cn.com/problems/subdomain-visit-count/</a>
 */
public class SubdomainVisits {

    public static void main(String[] args) {
        String[] cpdomains = {"9001 discuss.leetcode.com"};

        SubdomainVisits executor = new SubdomainVisits();

        List<String> strings = executor.subdomainVisits(cpdomains);

        CommonUtil.printStringList(strings);
    }

    private List<String> subdomainVisits(String[] cpdomains) {
        Map<String, Integer> domainCount = new HashMap<>();

        for (String domain : cpdomains) {
            // handle domain with given rule
            // checkout count
            int count = Integer.parseInt(domain.split(" ")[0]);
            String sub = domain.split(" ")[1];

            while (sub.contains(".")) {
                domainCount.put(sub, domainCount.getOrDefault(sub, 0) + count);

                sub = sub.substring(sub.indexOf(".") + 1);
            }

            // handle the last or the domain has no '.'
            domainCount.put(sub, domainCount.getOrDefault(sub, 0) + count);
        }

        List<String> res = new LinkedList<>();
        domainCount.forEach((domain, count) -> {
            res.add(count + " " + domain);
        });

        return res;
    }

}
