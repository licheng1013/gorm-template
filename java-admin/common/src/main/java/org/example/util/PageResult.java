package org.example.util;

import lombok.Data;

/**
 * @author lc
 * @since 2023/7/5
 */
@Data
public class PageResult {
    private Object list;
    private long total;

    public PageResult(Object list, long total) {
        this.list = list;
        this.total = total;
    }
}
