# Copyright 2022 Criticality Score Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.24.2@sha256:1ecc479bc712a6bdb56df3e346e33edcc141f469f82840bab9f4bc2bc41bf91d AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

FROM base AS criticality_score
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 go build ./cmd/criticality_score
RUN chmod -R 0775 /src/config/scorer/*

FROM gcr.io/distroless/base:nonroot@sha256:fa5f94fa433728f8df3f63363ffc8dec4adcfb57e4d8c18b44bceccfea095ebc
COPY --from=criticality_score /src/criticality_score ./criticality_score
COPY --from=criticality_score /src/config/scorer/* ./config/scorer/
ENTRYPOINT ["./criticality_score"]
