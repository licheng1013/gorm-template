package org.example.exception;

import lombok.extern.slf4j.Slf4j;
import org.example.util.HttpServletUtil;
import org.example.util.R;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

/**
 * 异常处理
 * @author lc
 */
@RestControllerAdvice
@Slf4j
public class GlobalExceptionHandler {

    /**
     * 处理 ServiceException 异常
     */
    @ExceptionHandler(ServiceException.class)
    public R<String> doHandleServiceException(ServiceException e) {
        // 打印请求参数
        log.info("请求参数: {}", HttpServletUtil.getParamMap());
        e.printStackTrace();
        return R.fail(e.getMessage());
    }

    /**
     * 处理 Exception 异常
     */
    @ExceptionHandler(Exception.class)
    public R<String> doHandleServiceException(Exception e) {
        e.printStackTrace();
        return R.fail("系统繁忙");
    }
}
