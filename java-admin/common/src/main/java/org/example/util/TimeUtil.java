package org.example.util;

import cn.hutool.core.date.DateUtil;

/**
 * Time Util
 * @author lc
 * @since 2023/7/5
 */
public class TimeUtil {
    // Get the start date string for today.
    public static String getTodayStartStr() {
       return DateUtil.beginOfDay(DateUtil.date()).toString();
    }
    // Get the end date string for today.
    public static String getTodayEndStr() {
        return DateUtil.endOfDay(DateUtil.date()).toString();
    }

    // Get the start date string for this month.
    public static String getMonthStartStr() {
        return DateUtil.beginOfMonth(DateUtil.date()).toString();
    }

    // Get the end date string for this month.
    public static String getMonthEndStr() {
        return DateUtil.endOfMonth(DateUtil.date()).toString();
    }

    // Get the start date string for this year.
    public static String getYearStartStr() {
        return DateUtil.beginOfYear(DateUtil.date()).toString();
    }

    // Get the end date string for this year.
    public static String getYearEndStr() {
        return DateUtil.endOfYear(DateUtil.date()).toString();
    }

    // Get the start date string for this week.
    public static String getWeekStartStr() {
        return DateUtil.beginOfWeek(DateUtil.date()).toString();
    }

    // Get the end date string for this week.
    public static String getWeekEndStr() {
        return DateUtil.endOfWeek(DateUtil.date()).toString();
    }

    // Test the function.
    public static void main(String[] args) {
        System.out.println(getTodayStartStr());
        System.out.println(getTodayEndStr());
        System.out.println(getMonthStartStr());
        System.out.println(getMonthEndStr());
        System.out.println(getYearStartStr());
        System.out.println(getYearEndStr());
        System.out.println(getWeekStartStr());
        System.out.println(getWeekEndStr());
    }

}
