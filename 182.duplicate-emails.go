//182. 查找重复的电子邮箱
//https://leetcode-cn.com/problems/duplicate-emails/
//# Write your MySQL query statement below
//--解法1
//select Email from Person group by Email having count(Email) > 1;

//--解法2
//select Email from (select count(1) as t,Email from Person group by Email)r  where r.t>1;

//--解法3
//select distinct(p1.Email) from Person p1  join Person  p2 on p1.Email = p2.Email AND p1.Id!=p2.Id