// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetKubernetesClusterAutoUpgrade {
    /**
     * @return True if Kubernetes patch version auto upgrades is enabled.
     * 
     */
    private Boolean enable;
    /**
     * @return The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
     * 
     */
    private String maintenanceWindowDay;
    /**
     * @return The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
     * 
     */
    private Integer maintenanceWindowStartHour;

    private GetKubernetesClusterAutoUpgrade() {}
    /**
     * @return True if Kubernetes patch version auto upgrades is enabled.
     * 
     */
    public Boolean enable() {
        return this.enable;
    }
    /**
     * @return The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
     * 
     */
    public String maintenanceWindowDay() {
        return this.maintenanceWindowDay;
    }
    /**
     * @return The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
     * 
     */
    public Integer maintenanceWindowStartHour() {
        return this.maintenanceWindowStartHour;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubernetesClusterAutoUpgrade defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enable;
        private String maintenanceWindowDay;
        private Integer maintenanceWindowStartHour;
        public Builder() {}
        public Builder(GetKubernetesClusterAutoUpgrade defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enable = defaults.enable;
    	      this.maintenanceWindowDay = defaults.maintenanceWindowDay;
    	      this.maintenanceWindowStartHour = defaults.maintenanceWindowStartHour;
        }

        @CustomType.Setter
        public Builder enable(Boolean enable) {
            this.enable = Objects.requireNonNull(enable);
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceWindowDay(String maintenanceWindowDay) {
            this.maintenanceWindowDay = Objects.requireNonNull(maintenanceWindowDay);
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceWindowStartHour(Integer maintenanceWindowStartHour) {
            this.maintenanceWindowStartHour = Objects.requireNonNull(maintenanceWindowStartHour);
            return this;
        }
        public GetKubernetesClusterAutoUpgrade build() {
            final var o = new GetKubernetesClusterAutoUpgrade();
            o.enable = enable;
            o.maintenanceWindowDay = maintenanceWindowDay;
            o.maintenanceWindowStartHour = maintenanceWindowStartHour;
            return o;
        }
    }
}
