# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class BrewInstallerName < Formula
  desc "A command-line tool to select default applications, based on duti"
  homepage "https://github.com/tsonglew/dutis"
  version "0.1.1"
  license "MIT"

  depends_on "git"
  depends_on "go"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/tsonglew/dutis/releases/download/v0.1.1/dutis_Darwin_arm64.tar.gz"
      sha256 "6608f18925e56f032c7ee619fc9a67362ae76b4b13fdbc66f09747c2b6411b5d"

      def install
        bin.install "dutis"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/tsonglew/dutis/releases/download/v0.1.1/dutis_Darwin_x86_64.tar.gz"
      sha256 "226d2ef19508766b3b5f2f878b27065133ff39ce745b6c3cecc6c5feaf333c0f"

      def install
        bin.install "dutis"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/tsonglew/dutis/releases/download/v0.1.1/dutis_Linux_arm64.tar.gz"
      sha256 "d9fe10866b385af725e1da7e230d1c544f822202b0406684a7ecf95756808e37"

      def install
        bin.install "dutis"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/tsonglew/dutis/releases/download/v0.1.1/dutis_Linux_x86_64.tar.gz"
      sha256 "8e232f013d739fc92ca14f9851afe31aeb769e68e63ac1119360ebb596830944"

      def install
        bin.install "dutis"
      end
    end
  end
end
