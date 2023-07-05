package org.example.util;

import lombok.Data;

import java.util.List;

/**
 * @author lc
 * @since 2023/7/5
 */
@Data
public class PageDto {
    private Integer page;
    private Integer size;
    private List<String> ids;
}
