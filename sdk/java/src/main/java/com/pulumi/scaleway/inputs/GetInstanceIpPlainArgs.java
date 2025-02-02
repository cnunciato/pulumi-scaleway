// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceIpPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceIpPlainArgs Empty = new GetInstanceIpPlainArgs();

    /**
     * The IPv4 address to retrieve
     * Only one of `address` and `id` should be specified.
     * 
     */
    @Import(name="address")
    private @Nullable String address;

    /**
     * @return The IPv4 address to retrieve
     * Only one of `address` and `id` should be specified.
     * 
     */
    public Optional<String> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * The ID of the IP address to retrieve
     * Only one of `address` and `id` should be specified.
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return The ID of the IP address to retrieve
     * Only one of `address` and `id` should be specified.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    private GetInstanceIpPlainArgs() {}

    private GetInstanceIpPlainArgs(GetInstanceIpPlainArgs $) {
        this.address = $.address;
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceIpPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceIpPlainArgs $;

        public Builder() {
            $ = new GetInstanceIpPlainArgs();
        }

        public Builder(GetInstanceIpPlainArgs defaults) {
            $ = new GetInstanceIpPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The IPv4 address to retrieve
         * Only one of `address` and `id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable String address) {
            $.address = address;
            return this;
        }

        /**
         * @param id The ID of the IP address to retrieve
         * Only one of `address` and `id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        public GetInstanceIpPlainArgs build() {
            return $;
        }
    }

}
