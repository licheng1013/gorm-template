package org.example.middleware;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

/**
 * @author lc
 */
@Configuration
public class AuthorizeConfig implements WebMvcConfigurer  {
    @Autowired
    private InterceptorConfig interceptorConfig;
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        //拦截所有路径
        registry.addInterceptor(interceptorConfig).addPathPatterns("/**");
    }
}
