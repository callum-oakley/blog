```
template = "post"
title = "Network-wide ad blocking with Pi-hole"
date = "6th of June 2020"
```

Recently I read [Against an Increasingly User-Hostile Web][] by Parimal Satyal (which is a few years old now, but increasingly relevant). It’s a brilliant piece, and if you can only read one thing today, read that instead of this. I left the article pretty upset, and, in need of *somewhere* to channel that energy, I set out to install a [Pi-hole][] on my home network.

Pi-hole is a network-wide advert and tracking blocker, which you can run on a [Raspberry Pi][]. (I’m more concerned with tracking than adverts, but please consider “ads” to be an abbreviation for “adverts and tracking” for the rest of this article.) Pi-hole blocks ads on every browser and app on every device on your local network, without you having to do any configuration on device. Somewhere you can’t normally install an ad blocker? No ads. A friend visits and connects to your wifi? No ads for them either.

All this is possible because ads are very often served from a different domain than the content you actually want to load. The Pi-hole then, poses as a <span class=sc>dns</span> server (responsible for mapping domain names to <span class=sc>ip</span> addresses) and refuses to resolve domains that it knows host ads – forwarding everything else to a real <span class=sc>dns</span> server of your choice. The result is that adverts never even have a chance to load, usually leaving a calming empty space where they would have been, and that the Googles and the Facebooks of the world can no longer follow your every move around the web. The “no configuration on your devices” magic is achieved by configuring your router to use the Pi-hole as its <span class=sc>dns</span> server, or by using the Pi-hole’s built in <span class=sc>dhcp</span> server (more on that later).

The setup was more straightforward than I expected it to be, and if you want to install one yourself I recommend primarily following the [Raspberry Pi set up guide][] and then the [Pi-hole docs][Pi-hole], but partly for my own reference, and partly because someone out there might find it useful, here are the steps I went through:

## Ingredients

- [Raspberry Pi Zero W][]
- <span class=sc>sd</span> card with [<span class=sc>noobs</span>][] (you can often buy these pre-loaded from the same place you buy the Pi)
- [Power supply][]
- Micro <span class=sc>usb</span> to <span class=sc>usb-a</span> adapter
- <span class=sc>usb</span> keyboard
- Mini <span class=sc>hdmi</span> to full sized <span class=sc>hdmi</span> adapter
- <span class=sc>hdmi</span> monitor

I didn’t have a <span class=sc>usb</span> mouse, and the Pi Zero only has one <span class=sc>usb</span> port anyway, but thankfully the <span class=sc>noobs</span> installer is very easy to run through with only a keyboard.

## Basic Raspberry Pi setup

1. Insert the <span class=sc>sd</span> card into your Pi, plug in your keyboard and monitor, and only then hook it up to the power supply.
2. You should be greeted with the <span class=sc>noobs</span> installer. Connect to wifi and then follow the prompts to install Raspbian (or maybe Raspberry <span class=sc>os</span> now since the name changed recently).
3. Follow the post install guide that pops up when you arrive at the desktop for the first time to configure language, wifi, etc.

## Remote access

I only have one monitor, which I need for my computer, so it was a priority for me to get <span class=sc>ssh</span> access to the Pi as soon as possible.

1. Find the <span class=sc>ip</span> address of your Pi and make a note of it – we’ll need it a few times below.
2. Enable <span class=sc>ssh</span> access with a password as described [here][ssh].
3. Copy a key over as described [here][passwordless-ssh].

Now you can unplug the monitor and keyboard, and do everything else over <span class=sc>ssh</span>.

## Install Pi-hole

Once you’ve got <span class=sc>ssh</span> access to your Pi, you can install Pi-hole by [piping the install script into bash][pi-hole install] (there are other options if you find piping to bash objectionable) and following some more prompts. The defaults all looked good to me.

## Route your network’s <span class=cc>dns</span> traffic to the Pi-hole

If your router supports it, it looks like the easiest final step is to [set your Pi-hole’s <span class=sc>ip</span> address to be your router’s *only* <span class=sc>dns</span> entry][router dns]. Unfortunately, my router doesn’t let me change the <span class=sc>dns</span> entries. Instead, I had to disable <span class=sc>dhcp</span> on the router, and [enable Pi-hole’s built in <span class=sc>dhcp</span> server][pi-hole dhcp].

A <span class=sc>dhcp</span> server assigns <span class=sc>ip</span> addresses to devices on your network, as well as telling devices where to find the router and the <span class=sc>dns</span> server (among other things). By disabling the router’s <span class=sc>dhcp</span> server, devices will automatically use the Pi-hole’s instead, which also means the Pi-hole becomes the authority on which <span class=sc>dns</span> server to use, and can point to itself.

The configuration here is mostly straightforward but check the Pi’s <span class=sc>dhcp</span> settings carefully. I messed up and failed to set the router <span class=sc>ip</span> address (wrongly assuming the default was correct), which briefly left me able to resolve <span class=sc>dns</span> queries but unable to actually load any web pages… If you’re feeling patient then devices will move to the new <span class=sc>dns</span> server next time they renew their <span class=sc>dhcp</span> lease, but you can accelerate the process by finding the option to *Renew Lease* in your device’s network settings.

* * *

Since setting it up, the Pi-hole has been humming along quietly blocking an astonishing 23% of unwanted <span class=sc>dns</span> traffic with no issues, and the internet is a better place because of it.

One final note: if you are able, please find other ways to support content creators and services that are normally funded by ad revenue. Buy a digital subscription to your favourite newspaper, pay for services and apps you like, support creators directly. I would love to see the industry at large move away from aggressive tracking and advertising, but the only way we’re going to see that happen is if we make alternatives viable.

[Against an Increasingly User-Hostile Web]: https://neustadt.fr/essays/against-a-user-hostile-web/
[Pi-hole]: https://pi-hole.net
[Raspberry Pi]: https://www.raspberrypi.org
[Raspberry Pi set up guide]: https://projects.raspberrypi.org/en/projects/raspberry-pi-setting-up
[Raspberry Pi Zero W]: https://www.raspberrypi.org/products/raspberry-pi-zero-w/
[<span class=sc>noobs</span>]: https://www.raspberrypi.org/downloads/noobs/
[power supply]: https://www.raspberrypi.org/products/raspberry-pi-universal-power-supply/
[ssh]: https://www.raspberrypi.org/documentation/remote-access/ssh/
[passwordless-ssh]: https://www.raspberrypi.org/documentation/remote-access/ssh/passwordless.md
[pi-hole install]: https://github.com/pi-hole/pi-hole/#one-step-automated-install
[router dns]: https://discourse.pi-hole.net/t/how-do-i-configure-my-devices-to-use-pi-hole-as-their-dns-server/245
[pi-hole dhcp]: https://discourse.pi-hole.net/t/how-do-i-use-pi-holes-built-in-dhcp-server-and-why-would-i-want-to/3026
