package org.example.util;

import lombok.Data;

/**
 * 泛型支持
 * @author lc
 * @since  2020/11/30
 */
@Data
public class R<T> {
    private Object msg;
    private int code ;
    private T data;
    public static <T> R<T> okMsg(T msg){
        R<T> jsonResult = new R<>();
        jsonResult.setMsg(msg);
        jsonResult.setCode(0);
        return jsonResult;
    }
    public static <T> R<T> okData(T data){
        R<T> jsonResult = new R<>();
        jsonResult.setData(data);
        jsonResult.setCode(0);
        return jsonResult;
    }
    public static <T> R<T> fail(T msg){
        R<T> jsonResult = new R<>();
        jsonResult.setMsg(msg);
        jsonResult.setCode(-1);
        return jsonResult;
    }
}
