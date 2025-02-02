// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionCronArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionCronArgs Empty = new FunctionCronArgs();

    /**
     * The key-value mapping to define arguments that will be passed to your function’s event object
     * during
     * 
     */
    @Import(name="args", required=true)
    private Output<String> args;

    /**
     * @return The key-value mapping to define arguments that will be passed to your function’s event object
     * during
     * 
     */
    public Output<String> args() {
        return this.args;
    }

    /**
     * The function ID to link with your cron.
     * 
     */
    @Import(name="functionId", required=true)
    private Output<String> functionId;

    /**
     * @return The function ID to link with your cron.
     * 
     */
    public Output<String> functionId() {
        return this.functionId;
    }

    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return (Defaults to provider `region`) The region
     * in where the job was created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     * 
     */
    @Import(name="schedule", required=true)
    private Output<String> schedule;

    /**
     * @return Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     * 
     */
    public Output<String> schedule() {
        return this.schedule;
    }

    private FunctionCronArgs() {}

    private FunctionCronArgs(FunctionCronArgs $) {
        this.args = $.args;
        this.functionId = $.functionId;
        this.region = $.region;
        this.schedule = $.schedule;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionCronArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionCronArgs $;

        public Builder() {
            $ = new FunctionCronArgs();
        }

        public Builder(FunctionCronArgs defaults) {
            $ = new FunctionCronArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param args The key-value mapping to define arguments that will be passed to your function’s event object
         * during
         * 
         * @return builder
         * 
         */
        public Builder args(Output<String> args) {
            $.args = args;
            return this;
        }

        /**
         * @param args The key-value mapping to define arguments that will be passed to your function’s event object
         * during
         * 
         * @return builder
         * 
         */
        public Builder args(String args) {
            return args(Output.of(args));
        }

        /**
         * @param functionId The function ID to link with your cron.
         * 
         * @return builder
         * 
         */
        public Builder functionId(Output<String> functionId) {
            $.functionId = functionId;
            return this;
        }

        /**
         * @param functionId The function ID to link with your cron.
         * 
         * @return builder
         * 
         */
        public Builder functionId(String functionId) {
            return functionId(Output.of(functionId));
        }

        /**
         * @param region (Defaults to provider `region`) The region
         * in where the job was created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region (Defaults to provider `region`) The region
         * in where the job was created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param schedule Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
         * executed.
         * 
         * @return builder
         * 
         */
        public Builder schedule(Output<String> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
         * executed.
         * 
         * @return builder
         * 
         */
        public Builder schedule(String schedule) {
            return schedule(Output.of(schedule));
        }

        public FunctionCronArgs build() {
            $.args = Objects.requireNonNull($.args, "expected parameter 'args' to be non-null");
            $.functionId = Objects.requireNonNull($.functionId, "expected parameter 'functionId' to be non-null");
            $.schedule = Objects.requireNonNull($.schedule, "expected parameter 'schedule' to be non-null");
            return $;
        }
    }

}
