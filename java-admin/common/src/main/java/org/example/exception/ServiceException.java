package org.example.exception;

import lombok.Getter;

/**
 * 自定义异常
 *
 * @author lc
 */
@Getter
public class ServiceException extends RuntimeException {
    private final int code;
    public ServiceException(String message) {
        this(message, -1);
    }
    public ServiceException(String message, int code) {
        super(message);
        this.code = code;
    }
}

