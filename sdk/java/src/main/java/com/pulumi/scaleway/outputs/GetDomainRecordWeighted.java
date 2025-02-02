// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDomainRecordWeighted {
    private String ip;
    private Integer weight;

    private GetDomainRecordWeighted() {}
    public String ip() {
        return this.ip;
    }
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainRecordWeighted defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ip;
        private Integer weight;
        public Builder() {}
        public Builder(GetDomainRecordWeighted defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetDomainRecordWeighted build() {
            final var o = new GetDomainRecordWeighted();
            o.ip = ip;
            o.weight = weight;
            return o;
        }
    }
}
