// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcPublicGatewayDhcpReservationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcPublicGatewayDhcpReservationPlainArgs Empty = new GetVpcPublicGatewayDhcpReservationPlainArgs();

    /**
     * The MAC address of the reservation to retrieve
     * 
     */
    @Import(name="macAddress")
    private @Nullable String macAddress;

    /**
     * @return The MAC address of the reservation to retrieve
     * 
     */
    public Optional<String> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }

    /**
     * The ID of the Reservation to retrieve
     * 
     */
    @Import(name="reservationId")
    private @Nullable String reservationId;

    /**
     * @return The ID of the Reservation to retrieve
     * 
     */
    public Optional<String> reservationId() {
        return Optional.ofNullable(this.reservationId);
    }

    /**
     * `zone`) The zone in which
     * the image exists.
     * 
     */
    @Import(name="zone")
    private @Nullable String zone;

    /**
     * @return `zone`) The zone in which
     * the image exists.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetVpcPublicGatewayDhcpReservationPlainArgs() {}

    private GetVpcPublicGatewayDhcpReservationPlainArgs(GetVpcPublicGatewayDhcpReservationPlainArgs $) {
        this.macAddress = $.macAddress;
        this.reservationId = $.reservationId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcPublicGatewayDhcpReservationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcPublicGatewayDhcpReservationPlainArgs $;

        public Builder() {
            $ = new GetVpcPublicGatewayDhcpReservationPlainArgs();
        }

        public Builder(GetVpcPublicGatewayDhcpReservationPlainArgs defaults) {
            $ = new GetVpcPublicGatewayDhcpReservationPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param macAddress The MAC address of the reservation to retrieve
         * 
         * @return builder
         * 
         */
        public Builder macAddress(@Nullable String macAddress) {
            $.macAddress = macAddress;
            return this;
        }

        /**
         * @param reservationId The ID of the Reservation to retrieve
         * 
         * @return builder
         * 
         */
        public Builder reservationId(@Nullable String reservationId) {
            $.reservationId = reservationId;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which
         * the image exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable String zone) {
            $.zone = zone;
            return this;
        }

        public GetVpcPublicGatewayDhcpReservationPlainArgs build() {
            return $;
        }
    }

}
