package org.example.exception;

/**
 * 自定义异常
 * @author lc
 */
public class ServiceException extends RuntimeException{
    public ServiceException(String message) {
        super(message);
    }
}

