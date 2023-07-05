package org.example.middleware;


import lombok.extern.slf4j.Slf4j;
import org.example.util.HttpServletUtil;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * @author lc
 */
@Configuration
@Slf4j
public class InterceptorConfig implements HandlerInterceptor {

    public static final String Authorize = "Authorization";

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
        //排除静态资源因为SpringBoot能够访问static下的文件
        if (!(handler instanceof HandlerMethod handlerMethod)){
            return true;
        }
        PassToken passToken = handlerMethod.getMethodAnnotation(PassToken.class);
        if (passToken != null) {
            return true;
        }
        // TODO 你的用户处理逻辑！ 例如获取Jwt的Token 处理认证等业务
        String token = HttpServletUtil.getHeader(Authorize);
        log.info("token: {}", token);
        // 处理完成后返回 true 即可继续执行业务, false 则中止业务！
        return true;
    }
}
