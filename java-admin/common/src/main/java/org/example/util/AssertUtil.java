package org.example.util;

import cn.hutool.core.util.StrUtil;
import org.example.exception.ServiceException;

/**
 * @author lc
 * @since 2023/7/5
 */
public class AssertUtil {
    // Is Empty
    public static void AssertEmpty(String str, String message) {
        if (StrUtil.isBlankIfStr(str)) {
            throw new ServiceException(message);
        }
    }

    // Is Empty and Greater than Length
    public static void AssertEmptyGtLength(String str, int length, String message) {
        if (StrUtil.isBlankIfStr(str) || str.length() > length) {
            throw new ServiceException(message);
        }
    }

    // Is Empty and Greater not eq Length
    public static void AssertEmptyGtNotEqLength(String str, int length, String message) {
        if (StrUtil.isBlankIfStr(str) || str.length() != length) {
            throw new ServiceException(message);
        }
    }

    // Is True
    public static void AssertTrue(boolean bool, String message) {
        if (bool) {
            throw new ServiceException(message);
        }
    }

    // Assert False
    public static void AssertFalse(boolean bool, String message) {
       AssertTrue(!bool, message);
    }

    // Assert Number eq Zero and Null
    public static void AssertNumberEqZeroAndNull(Integer num, String message) {
        if (num == null || num == 0) {
            throw new ServiceException(message);
        }
    }

    // Assert Number of Range
    public static void AssertNumberRange(Integer num, Integer min, Integer max, String message) {
        if (num == null || num < min || num > max) {
            throw new ServiceException(message);
        }
    }

    public static void AssertError(String error) {
        throw new ServiceException(error);
    }
}
