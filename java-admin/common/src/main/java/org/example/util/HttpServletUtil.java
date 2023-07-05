package org.example.util;

import cn.hutool.extra.servlet.ServletUtil;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import java.util.Map;
import java.util.Objects;

/**
 * 必须在web环境使用,否则会出现空指针异常
 * @author lc
 * @since 2023/7/5
 */
public final class HttpServletUtil {
    private HttpServletUtil() {
    }
    /**
     * 当前请求的路径
     */
    public static String getPath() {
        return getHttpServletRequest().getRequestURI();
    }
    /**
     * 获取HttpServletRequest
     */
    private static HttpServletRequest getHttpServletRequest() {
        return ((ServletRequestAttributes) Objects.requireNonNull(RequestContextHolder.getRequestAttributes())).getRequest();
    }
    /**
     * 获取所有请求参数
     */
    public static Map<String, String> getParamMap() {
        return ServletUtil.getParamMap(getHttpServletRequest());
    }
    /**
     * 获取某个请求头
     */
    public static String getHeader(String name) {
        return getHttpServletRequest().getHeader(name);
    }
    /**
     * 获取请求类型GET,POST
     */
    public static String getMethod() {
        return getHttpServletRequest().getMethod();
    }

}
